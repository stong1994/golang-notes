package main

func main() {
	for i := 1; i <= 30; i++ {
		println(getFib(i))
	}
}

func getFib(i int) int {
	if i <= 1 {
		return i
	}
	val := getFib(i-1) + getFib(i-2)
	return val
}
