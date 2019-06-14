package main

func main() {
	fib := getFib()
	for i := 0; i < 30; i++ {
		println(fib())
	}
}

func getFib() func() int {
	m := 0
	n := 1
	return func() int {
		fib := n
		m, n = n, m+n
		return fib
	}
}
