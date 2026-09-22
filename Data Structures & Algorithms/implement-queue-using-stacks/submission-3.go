type MyQueue struct {
	Inbox []int
	Outbox []int
}

func Constructor() MyQueue {
	return MyQueue{
		Inbox: []int{},
		Outbox: []int{},
	}
}

func (this *MyQueue) FromInToOut() {
	for len(this.Inbox) > 0 {
		value := this.Inbox[len(this.Inbox) - 1]
		this.Inbox = this.Inbox[:len(this.Inbox) - 1]
		this.Outbox = append(this.Outbox, value)
	}
}

func (this *MyQueue) Push(x int) {
	this.Inbox = append(this.Inbox, x)
}

func (this *MyQueue) Pop() int {
	if len(this.Outbox) == 0 {
		this.FromInToOut()
	}

	value := this.Outbox[len(this.Outbox) - 1]	
	this.Outbox = this.Outbox[:len(this.Outbox) - 1]

	return value
}

func (this *MyQueue) Peek() int {
	if len(this.Outbox) == 0 {
		this.FromInToOut()
	}

	return this.Outbox[len(this.Outbox) - 1]
}

func (this *MyQueue) Empty() bool {
	return len(this.Inbox) == 0 && len(this.Outbox) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
