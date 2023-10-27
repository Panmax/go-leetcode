package review_20231027

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    nil,
		minStack: []int{math.MaxInt32},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
