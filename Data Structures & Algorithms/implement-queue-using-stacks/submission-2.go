type Node struct {
	Value int
	Prev *Node
	Next *Node
}

type MyQueue struct {
	Head *Node
	Tail *Node
	Size int
}

func Constructor() MyQueue {
	return MyQueue{} 
}

func (this *MyQueue) Push(x int) {
	node := &Node{Value: x}

	if this.Size == 0 {
		this.Head = node
		this.Tail = node
	} else {
		this.Tail.Prev = node
		node.Next = this.Tail
		this.Tail = node
	}
	
	this.Size++
}

func (this *MyQueue) Pop() int {
	if this.Size == 0 {return -1}

	value := this.Head.Value

	if this.Size == 1 {
		this.Head = nil
		this.Tail = nil
		this.Size = 0
	} else {
		this.Head = this.Head.Prev
		this.Head.Next = nil
		this.Size--
	}

	return value
}

func (this *MyQueue) Peek() int {
	return this.Head.Value
}

func (this *MyQueue) Empty() bool {
	return this.Size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
