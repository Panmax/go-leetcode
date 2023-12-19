package review_20231219

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
	v := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return v
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
