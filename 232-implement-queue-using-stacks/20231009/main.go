package main

type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
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
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

func (this *MyQueue) in2out() {
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}
