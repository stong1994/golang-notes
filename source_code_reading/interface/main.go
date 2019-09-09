package main

type EmptyStruct struct {

}

func main() {
	var es interface{} = EmptyStruct{}
	var es2 interface{} = &EmptyStruct{}
	_, _ = es, es2
}