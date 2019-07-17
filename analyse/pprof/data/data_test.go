package data

import (
	"fmt"
	"testing"
)

const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	var data []string
	for i := 0; i < b.N; i++ {
		data = append(data, fmt.Sprintf("hello %d", i))
	}
}
