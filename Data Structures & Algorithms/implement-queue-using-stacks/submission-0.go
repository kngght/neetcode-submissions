type MyQueue struct {
	*list.List
}

func Constructor() MyQueue {
	return MyQueue{
		list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.PushBack(x)
}

func (this *MyQueue) Pop() int {
	front := this.Front()
	if front == nil {
		return 0
	}
	val := front.Value.(int)
	this.Remove(front)
	return val
}

func (this *MyQueue) Peek() int {
	front := this.Front()
	if front == nil {
		return 0
	}
	return front.Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.Len() == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
