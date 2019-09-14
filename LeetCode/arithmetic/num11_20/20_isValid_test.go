package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			"test1",
			"()",
			true,
		}, {
			"test2",
			"()[]{}",
			true,
		}, {
			"test3",
			"(]",
			false,
		}, {
			"test4",
			"([)]",
			false,
		}, {
			"test5",
			"{[]}",
			true,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := isValid(v.arg)
			if got != v.want {
				t.Errorf("want %t, got %t", v.want, got)
			}
		})
	}
}

// 栈
func my(s string) bool {
	var (
		ls = map[byte]byte{'(': ')', '[': ']', '{': '}'}
		rs = map[byte]byte{')': '(', ']': '[', '}': '{'}
		m  = map[byte]struct{}{')': struct{}{}, ']': struct{}{}, '}': struct{}{}, '(': struct{}{}, '[': struct{}{}, '{': struct{}{}}
	)

	arr := []byte{}
	for _, v := range []byte(s) {
		if _, ok := m[v]; ok {
			arr = append(arr, v)
		}
	}

	if len(arr) == 0 {
		return true
	}
	if len(arr)&1 == 1 {
		return false
	}

	stack := []byte{}
	for _, v := range arr {
		if _, ok := ls[v]; ok {
			stack = append(stack, v)
			continue
		}
		if _, ok := rs[v]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == rs[v] {
				stack = append(stack[:len(stack)-1])
				continue
			} else {
				return false
			}
		}

	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	matchs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack = make([]byte, 0, len(s))
	for _, v := range []byte(s) {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) != 0 && matchs[v] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
