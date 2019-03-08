package main

import (
	"fmt"
	_ "golang-learning/interview/Golang-2019-2-20/1-init-order/data"
	_ "golang-learning/interview/Golang-2019-2-20/1-init-order/info"
)

func init() {
	fmt.Println("main init function")
}

func main() {
	fmt.Println("main function")
}
