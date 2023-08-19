package minstack

import (
	"fmt"
	"testing"
)

func TestMinStack_Push(t *testing.T) {
	obj := Constructor()
	obj.Push(5)
	fmt.Println(obj.GetMin())
	obj.Push(8)
	fmt.Println(obj.GetMin())
	obj.Push(0)
	fmt.Println(obj.GetMin())
	obj.Push(1)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	if obj.GetMin() != 5 {
		t.Error("test case failed")
	}
}
