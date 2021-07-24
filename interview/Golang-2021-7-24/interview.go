package Golang_2021_7_24

// https://www.zhihu.com/question/60952598/answer/1996173175
// 1. 两个函数哪个效率更高
const matrixLength = 20000

func foo() {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			matrixA[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}
}

func bar() {
	matrixA := createMatrix(matrixLength)
	matrixB := createMatrix(matrixLength)

	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			matrixA[i][j] = matrixA[i][j] + matrixB[j][i]
		}
	}
}

// 2. 多核CPU是如何保持状态一致的

// 3. 运行结果
func q3() {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}

// 4. rune类型
// 5. 每个goroutine 分别打印cat dog fish，并按顺序打印。总共打印100次
// 6. 对channel、mutex、goroutine的了解
// 7. linux的信号
// 8. 内存分配机制

// 1. 考点：cpu缓存，空间的局部性。
// 5
func q5() {
	count := make(chan struct{})
	finish := make(chan struct{})
	cat := make(chan struct{})
	dog := make(chan struct{})
	fish := make(chan struct{})

	go func() {
		for {
			select {
			case <-finish:
				break
			case <-cat:
				fmt.Println("cat")
				dog <- struct{}{}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-finish:
				break
			case <-dog:
				fmt.Println("dog")
				fish <- struct{}{}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-finish:
				break
			case <-fish:
				fmt.Println("fish")
				count <- struct{}{}
			}
		}
	}()

	go func() {
		i := 0
		for {
			select {
			case <-count:
				i++
				fmt.Println(i)
				if i == 100 {
					finish <- struct{}{}
					return
				}
				cat <- struct{}{}
			}
		}
	}()

	cat <- struct{}{}
	<-finish
}
