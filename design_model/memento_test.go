package design_model

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	circle := &Circle{}
	circle.DrawColor("red")
	fmt.Println("当前颜色", circle.color)
	backup := circle.Save()
	fmt.Println("备份")

	circle.DrawColor("yellow")
	fmt.Println("当前颜色", circle.color)
	circle.Redo(backup)
	fmt.Println("恢复")
	fmt.Println("当前颜色", circle.color)
}
