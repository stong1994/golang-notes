package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

/*
在项目中我们常用到redis或者其它中间件来缓存数据，
然而这里通过反射直接缓存了方法。
还需要深入了解这方面知识。
*/
func main() {
	ch, err := Cacher(AddSlowly, 2*time.Second)
	if err != nil {
		panic(err)
	}
	chAddSlowly := ch.(func(int, int) int)
	for i := 0; i < 5; i++ {
		start := time.Now()
		result := chAddSlowly(1, 2)
		end := time.Now()
		fmt.Println("got result", result, "in", end.Sub(start))
	}
	time.Sleep(3 * time.Second)
	start := time.Now()
	result := chAddSlowly(1, 2)
	end := time.Now()
	fmt.Println("got result", result, "in", end.Sub(start))
}

func AddSlowly(a, b int) int {
	time.Sleep(100 * time.Millisecond)
	return a + b
}

type outExp struct {
	out    []reflect.Value
	expiry time.Time
}

func Cacher(f interface{}, expiration time.Duration) (interface{}, error) {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		return nil, errors.New("Only for functions")
	}

	inType, err := buildInStruct(ft)
	if err != nil {
		return nil, err
	}

	if ft.NumOut() == 0 {
		return nil, errors.New("Must have at least one returned value")
	}

	m := map[interface{}]outExp{}
	fv := reflect.ValueOf(f)
	cacher := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		iv := reflect.New(inType).Elem()
		for k, v := range args {
			iv.Field(k).Set(v)
		}
		ivv := iv.Interface()
		ov, ok := m[ivv]
		now := time.Now()
		// 如果不存在或者过期了，调用函数获取值，然后赋值。
		if !ok || ov.expiry.Before(now) {
			ov.out = fv.Call(args)
			ov.expiry = now.Add(expiration)
			m[ivv] = ov
		}
		return ov.out
	})
	return cacher.Interface(), nil
}

func buildInStruct(ft reflect.Type) (reflect.Type, error) {
	if ft.NumIn() == 0 {
		return nil, errors.New("Must have at least one param")
	}
	var sf []reflect.StructField
	for i := 0; i < ft.NumIn(); i++ {
		ct := ft.In(i)
		if !ct.Comparable() {
			return nil, fmt.Errorf("parameter %d of type %s and kind %v is not comparable", i+1, ct.Name(), ct.Kind())
		}
		sf = append(sf, reflect.StructField{
			Name: fmt.Sprintf("F%d", i),
			Type: ct,
		})
	}
	s := reflect.StructOf(sf)
	return s, nil
}
