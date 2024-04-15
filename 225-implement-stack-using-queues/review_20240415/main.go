package main

type MyStack struct {
	queue1, queue2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *MyStack) Pop() int {
	x := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
