package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	fly := NewFlyweight()
	animal := NewAnimal()
	fly.SetElement("base_animal", animal)
	// 生成两个动物
	cat := fly.GetElement("base_animal").(IAnimal)
	dog := fly.GetElement("base_animal").(IAnimal)

	// 成长
	cat = NewGrowUp(cat, 10, 1)
	dog = NewGrowUp(dog, 50, 1)
	// 获取体重和年龄
	catWeight := cat.GetWeight()
	catAge := cat.GetAge()
	dogWeight := dog.GetWeight()
	dogAge := dog.GetAge()

	fmt.Printf("cat weight: %d, cat age: %d\n", catWeight, catAge)
	fmt.Printf("dog weight: %d, dog age: %d\n", dogWeight, dogAge)
}

// BenchmarkFlyweight-8   	10000000	       127 ns/op	      82 B/op	       0 allocs/op
func BenchmarkFlyweight(b *testing.B) {
	b.ReportAllocs()
	fly := NewFlyweight()
	animal := NewAnimal()
	fly.SetElement("base_animal", animal)

	arr := []IAnimal{}

	for i := 0; i < b.N; i++ {
		arr = append(arr, fly.GetElement("base_animal").(IAnimal))

	}
}

// BenchmarkNewAnimal-8   	10000000	       135 ns/op	      98 B/op	       1 allocs/op
func BenchmarkNewAnimal(b *testing.B) {
	b.ReportAllocs()
	arr := []IAnimal{}
	for i:= 0; i< b.N; i++ {
		arr = append(arr, new(Animal))
	}
}