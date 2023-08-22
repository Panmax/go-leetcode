package main

import "fmt"

type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) in2out() {
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.out) <= 0 {
		this.in2out()
	}
	x := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.out) <= 0 {
		this.in2out()
	}
	x := this.out[len(this.out)-1]
	return x
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	fmt.Print(param_2)

	param_3 := obj.Peek()
	fmt.Print(param_3)

	param_4 := obj.Empty()
	fmt.Print(param_4)

}
