package main

func main() {
	m := 0
	n := 1
	for i := 0; i < 30; i++ {
		fib := n
		m, n = n, m+n
		println(fib)
	}
}
