package main

import "fmt"

// 快排的思想是在切片中找到一个元素a,然后令a左边的元素都小于a,a右边的元素都大于a,然后在左右两段切片中再找到这样的元素,以此类推,将会得到一个升序的切片
func main() {
	rowArr := []int{5, 3, 6, 8, 1, 3, 9, 0, 2, 10, 30, 18, 19, 4}
	quickSort2(rowArr, 0, len(rowArr)-1)
	fmt.Println(rowArr)

}

// 第一种方法
// arr为整个切片, l即left,是需要调整顺序的某段切片的第一个索引,r即right,与l相反是最后一位元素的索引
// 思路为对需要调整顺序的某段切片进行遍历,得到左右两个切片,然后对原切片进行赋值
func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	baseNum := arr[l]             // 以需要调整顺序的切片的第一个元素为基准
	leftArr := []int{}            // 小于基准的切片
	rightArr := []int{}           // 大于基准的切片
	for i := l + 1; i <= r; i++ { //
		if arr[i] > baseNum {
			rightArr = append(rightArr, arr[i])
		} else {
			leftArr = append(leftArr, arr[i])
		}
	}
	newArr := append(leftArr, baseNum)
	newArr = append(newArr, rightArr...)

	// 对需要调整顺序的一段切片的进行重新赋值
	for i := 0; i < len(newArr); i++ {
		arr[i+l] = newArr[i]
	}

	// 对基准左右两段切片分别进行调整顺序
	quickSort(arr, l, l+len(leftArr)-1)
	quickSort(arr, l+len(leftArr)+1, r)
}

// 第二种方法,快排的主要思路在在一段切片中,找到一个元素使得这段切片的左边都小于这个元素,右边都大于这个元素,可以将第一个元素作为基准元素,然后从最后一个元素开始比较,
// 如果大于基准元素,则比较倒数第二个元素,直到找到一个比基准元素小的,然后交换位置,然后从左边开始比较,一直找到比基准元素大的然后进行交换.
// 这样能够保证基准元素左边全是比它小的,而右边全是比它大的,且不需要创建新的切片,直接在原切片进行操作即可
func quickSort2(arr []int, l, r int) {
	if l >= r {
		return
	}
	m, n := l, r

	flag := l
	for l < r {
		for r > flag {
			if arr[r] < arr[flag] {
				arr[r], arr[flag] = arr[flag], arr[r]
				flag = r
				break
			}
			r--
		}
		for l < flag {
			if arr[l] > arr[flag] {
				arr[l], arr[flag] = arr[flag], arr[l]
				flag = l
				break
			}
			l++
		}
	}
	quickSort2(arr, m, flag-1)
	quickSort2(arr, flag+1, n)
}
