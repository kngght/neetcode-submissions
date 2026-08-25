type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	Cap        int
	Cache      map[int]*Node
	Head, Tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Cap:   capacity,
		Cache: make(map[int]*Node, capacity),
		Head:  head,
		Tail:  tail,
	}
}

func (c *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) insert(node *Node) {
	node.Prev = c.Head
	node.Next = c.Head.Next
	node.Next.Prev = node
	c.Head.Next = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.insert(node)
	return node.Val
}

func (this *LRUCache) Put(key int, val int) {
	if node, ok := this.Cache[key]; ok {
		node.Val = val
		this.remove(node)
		this.insert(node)
		return
	}

	node := &Node{
		Key: key,
		Val: val,
	}
	this.Cache[key] = node
	this.insert(node)
	if len(this.Cache) > this.Cap {
		delete(this.Cache, this.Tail.Prev.Key)
		this.remove(this.Tail.Prev)
	}
}
