type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{
		Queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyStack) Pop() int {
	if !this.Empty() {
		value := this.Queue[len(this.Queue) - 1]	
		this.Queue = this.Queue[:len(this.Queue) - 1]
		return value
	}
	return -1
}

func (this *MyStack) Top() int {
	if !this.Empty() {
		return this.Queue[len(this.Queue) - 1]	
	}
	return -1
}

func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
