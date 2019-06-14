package sort

import (
	"fmt"
	"golang-learning/tools/utils/rand_num"
	"testing"
)

func TestSort(t *testing.T) {
	type person struct {
		height float64
		weight float64
		age    uint8
	}
	lens := 10
	var persons = make([]*person, 0, 10) // 0-9的随机数
	for i := 0; i < lens; i++ {
		h := rand_num.RandNum(int64(lens))
		w := rand_num.RandNum(int64(lens))
		a := rand_num.RandNum(int64(lens))
		height := h*5 + 150 // 150 - 200
		weight := w*10 + 60 // 60 - 160
		age := a*2 + 10     // 10 - 30
		persons = append(persons, &person{height: float64(height), weight: float64(weight), age: uint8(age)})
	}
	// sort by height
	swap := func(i, j int) { persons[i], persons[j] = persons[j], persons[i] }
	less := func(i, j int) bool { return persons[i].height > persons[j].height }
	Sort(len(persons), swap, less)
	fmt.Println("sort by height")
	for _, v := range persons {
		fmt.Printf("height: %f, weight: %f, age: %d\n", v.height, v.weight, v.age)
	}
	// sort by weight
	Sort(len(persons),
		func(i, j int) {
			persons[i], persons[j] = persons[j], persons[i]
		},
		func(i, j int) bool {
			return persons[i].weight > persons[j].weight
		})
	fmt.Println("sort by weight")
	for _, v := range persons {
		fmt.Printf("height: %f, weight: %f, age: %d\n", v.height, v.weight, v.age)
	}
	// sort by age
	Sort(len(persons),
		func(i, j int) {
			persons[i], persons[j] = persons[j], persons[i]
		},
		func(i, j int) bool {
			return persons[i].age > persons[j].age
		})
	fmt.Println("sort by age")
	for _, v := range persons {
		fmt.Printf("height: %f, weight: %f, age: %d\n", v.height, v.weight, v.age)
	}
}
