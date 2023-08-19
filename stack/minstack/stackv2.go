package minstack

import (
	"fmt"
	"math"
)

type MinStack struct {
	top *node
	len int
	min int
}

type node struct {
	prev *node
	val  int
	min  int
}

func Constructor() MinStack {
	return MinStack{nil, 0, math.MinInt}
}

func (this *MinStack) Push(val int) {
	n := &node{this.top, val, val}
	this.top = n
	if this.len == 0 {
		this.min = n.min
	} else if this.min > n.min {
		this.min = n.min
	} else {
		n.min = this.min
	}
	this.len++
}

func (this *MinStack) Pop() {
	if this == nil || this.len == 0 {
		return
	}
	n := this.top
	this.top = n.prev
	if n.prev != nil {
		this.min = n.prev.min
	}
	this.len--
}

func (this *MinStack) Top() int {
	return this.top.val
}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) PrintStack() {
	n := this.top
	fmt.Println("Stack Started")
	for n != nil {
		fmt.Println(n.val)
		n = n.prev
	}
	fmt.Println("Stack Finished")
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
