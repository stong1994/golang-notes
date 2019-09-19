package main

import (
	"fmt"
	"strconv"
	"strings"
)
/*
2.Given a “flatten” dictionary object, whose keys are dot-separated.
For example, { ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3, ‘CC.D.E’: 4, ‘CC.D.F’: 5}.
Implement a function in any language to transform it to a “nested” dictionary object.
In the above case, the nested version is like:
{
  ‘A’: 1,
  ‘B’: {
    ‘A’: 2,
    ‘B’: 3,
  },
  ‘CC’: {
    ‘D’: {
      ‘E’: 4,
      ‘F’: 5,
    }
  }
}
It’s guaranteed that no keys in dictionary are prefixes of other keys.
 */
func main() {
	data := "{ 'A':1,'B.A':2,'B.B':3,'CC.D.E':4,'CC.D.F':5}"
	fmt.Println(dictionaries(data))
}
// TODO 忘记去掉单引号
func dictionaries(data string) map[string]interface{} {
	// 去掉大括号
	data = data[1 : len(data)-1]
	// 去掉空格
	data = strings.Replace(data, " ", "", -1)

	kvs := strings.Split(data, ",")

	res := make(map[string]interface{})

	for _, v := range kvs {
		// 根据冒号分割key 和value
		kv := strings.Split(v, ":")
		if len(kv) != 2 {
			panic(fmt.Sprintf("数据格式错误：%s", v))
		}
		value, err := strconv.ParseInt(kv[1], 10, 64)
		if err != nil {
			panic(err)
		}
		if !strings.Contains(kv[0], ".") {
			res[kv[0]] = value
			continue
		}
		ks := strings.Split(kv[0], ".")

		res[ks[0]] = keys(strings.Join(ks[1:], "."), value)
	}
	return res
}

func keys(key string, val int64) map[string]interface{} {
	m := make(map[string]interface{})
	if !strings.Contains(key, ".") {
		m[key] = val
		return m
	}
	ks := strings.Split(key, ".")
	k := ks[0]
	v := strings.Join(ks[1:], ".")

	m[k] = keys(v, val)
	return m
}

