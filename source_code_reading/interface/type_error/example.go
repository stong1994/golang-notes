package main

import (
	"fmt"
)

type MyError struct {
	Code int
	Msg  string
}

func (me MyError) Error() string {
	return me.Msg
}

func NilErr() *MyError {
	return nil
}

func main() {
	var err error
	err = NilErr()
	if err != nil {
		fmt.Println("has error")
		return
	}
	fmt.Println("no error")
}
