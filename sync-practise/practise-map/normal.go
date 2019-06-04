package main

// 在高并发下同时读写map会报错
func main() {
	m := map[int]int{}
	for i := 0; i < 40000; i++ {
		go write(m, i, i*i)
		go read(m, i)
	}
}

func read(m map[int]int, key int) int {
	return m[key]
}

func write(m map[int]int, key, val int) {
	m[key] = val
}
