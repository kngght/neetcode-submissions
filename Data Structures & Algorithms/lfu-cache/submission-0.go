type Node struct {
	Key, Val, Freq int
	Prev, Next *Node
}

type LFUCache struct {
	Cap int
	Cache map[int]*Node
	Head, Tail *Node
}

func Constructor(capacity int) LFUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Prev = head
	return LFUCache {
		Cap: capacity,
		Cache: make(map[int]*Node, capacity+1),
		Head: head,
		Tail: tail,
	} 
}

func (this *LFUCache) Remove() {
	if this.Tail.Prev == this.Head {
		return
	}
	delete(this.Cache, this.Tail.Prev.Key)
	this.Tail.Prev = this.Tail.Prev.Prev
	this.Tail.Prev.Next = this.Tail
}

func (this *LFUCache) Insert(node *Node) {
	curr := this.Tail
	for curr.Prev != nil && curr.Prev != this.Head && curr.Prev.Freq <= node.Freq {
		curr = curr.Prev
	}
	node.Prev = curr.Prev
	curr.Prev.Next = node
	curr.Prev = node
	node.Next = curr
}

func (this *LFUCache) Increase(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Freq++
	this.Insert(node)
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.Increase(node)
	return node.Val
}

func (this *LFUCache) Put(key int, value int)  {
	if node, ok := this.Cache[key]; ok {
		node.Val = value
		this.Increase(node)
		return
	}
	node := &Node{
		Key: key,
		Val: value,
		Freq: 1,
	}
	if len(this.Cache) >= this.Cap {
		this.Remove()
	}
	this.Cache[key] = node
	this.Insert(node)
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */
