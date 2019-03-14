package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	greeting := "hello"
	f := Foo{A: 10, B: "Salutations"}

	gVal := reflect.ValueOf(greeting)
	// not a pointer so all we can do is read it
	fmt.Println(gVal.Interface())

	gpVal := reflect.ValueOf(&greeting)
	// it’s a pointer, so we can change it, and it changes the underlying variable
	gpVal.Elem().SetString("goodbye")
	fmt.Println(greeting)

	// 改变原结构体中的值
	b := reflect.ValueOf(&f.A)
	b.Elem().SetInt(20)
	fmt.Printf("%+v, %d, %s\n", f, f.A, f.B)

	// 根据获取到的类型新建一个该类型的实例，并赋值
	fType := reflect.TypeOf(f)
	fVal := reflect.New(fType)
	fVal.Elem().Field(0).SetInt(30)
	fVal.Elem().Field(1).SetString("Greetings")
	f2 := fVal.Elem().Interface().(Foo)
	fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)
}
