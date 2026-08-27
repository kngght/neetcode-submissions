type Node struct {
	Val int
	Next *Node
}

type MyCircularQueue struct {
	FreeSpace int
	Head, Tail *Node
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue {
		FreeSpace: k,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {return false}
	if this.Tail != nil {
		this.Tail.Next = &Node{
			Val: value,
		}
		this.Tail = this.Tail.Next
	} else {
		this.Tail = &Node{Val: value}	
		this.Head = this.Tail
	}
	this.FreeSpace--
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {return false}
	if this.Head == this.Tail {
		this.Head = nil
		this.Tail = this.Head
	} else {
		this.Head = this.Head.Next
	}
	this.FreeSpace++
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {return -1}
	return this.Head.Val
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {return -1}
	return this.Tail.Val
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == nil 
}


func (this *MyCircularQueue) IsFull() bool {
	return this.FreeSpace == 0
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 