type Node struct {
    Key, Val, Freq int
    Prev, Next *Node
}

type LinkedList struct {
    Head, Tail *Node
    Size int
}

func NewLinkedList() *LinkedList {
    head, tail := &Node{}, &Node{}
    head.Next = tail
    tail.Prev = head
    return &LinkedList{
        Head: head,
        Tail: tail,
    }
}

func (l *LinkedList) AddAtTail(node *Node) {
    node.Prev = l.Tail.Prev
    node.Prev.Next = node
    node.Next = l.Tail
    l.Tail.Prev = node
    if l.Size == 0 {
        l.Head.Next = node 
    }
    l.Size++
}

func (l *LinkedList) Remove(node *Node) {
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
    node.Prev = nil
    node.Next = nil
    l.Size--
}

func (l *LinkedList) DeleteAtHead() *Node {
    if l == nil || l.Size == 0 {
        return nil 
    }
    node := l.Head.Next
    l.Remove(node)
    return node 
}

type LFUCache struct {
    Cap int
    MinFreq int
    Cache map[int]*Node
    Set map[int]*LinkedList
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        Cap: capacity,
        MinFreq: 0,
        Cache: make(map[int]*Node, capacity+1),
        Set: make(map[int]*LinkedList),
    }
}

func (this *LFUCache) Get(key int) int {
    node, ok := this.Cache[key]
    if !ok {
        return -1
    }
    oldFreq := node.Freq 
    oldList := this.Set[oldFreq]
    oldList.Remove(node)
    if oldFreq == this.MinFreq && oldList.Size == 0 {
        this.MinFreq++
    }
    node.Freq++
    newList, ok := this.Set[node.Freq]
    if !ok {
        newList = NewLinkedList()
        this.Set[node.Freq] = newList
    }
    newList.AddAtTail(node)
    return node.Val
}

func (this *LFUCache) Put(key int, value int)  {
    if this.Cap == 0 {
        return
    }
    node, ok := this.Cache[key]    
    if ok {
        node.Val = value  
        this.Get(key)
        return
    }
    if len(this.Cache) >= this.Cap {
        list := this.Set[this.MinFreq]
        deletedNode := list.DeleteAtHead()
        delete(this.Cache, deletedNode.Key)
        if list.Size == 0 {
            delete(this.Set, this.MinFreq)
        }
    }
    node = &Node{
        Key: key,
        Val: value,
        Freq: 1,
    }
    this.Cache[key] = node
    this.MinFreq = 1
    list, ok := this.Set[node.Freq]
    if !ok {
        list = NewLinkedList()
        this.Set[node.Freq] = list
    }
    list.AddAtTail(node)
}

// type Node struct {
// 	Key, Val, Freq int
// 	Prev, Next *Node
// }

// type LFUCache struct {
// 	Cap int
// 	Cache map[int]*Node
// 	Head, Tail *Node
// }

// func Constructor(capacity int) LFUCache {
// 	head, tail := &Node{}, &Node{}
// 	head.Next = tail
// 	tail.Prev = head
// 	return LFUCache {
// 		Cap: capacity,
// 		Cache: make(map[int]*Node, capacity+1),
// 		Head: head,
// 		Tail: tail,
// 	} 
// }

// func (this *LFUCache) Remove() {
// 	if this.Tail.Prev == this.Head {
// 		return
// 	}
// 	delete(this.Cache, this.Tail.Prev.Key)
// 	this.Tail.Prev = this.Tail.Prev.Prev
// 	this.Tail.Prev.Next = this.Tail
// }

// func (this *LFUCache) Insert(node *Node) {
// 	curr := this.Tail
// 	for curr.Prev != nil && curr.Prev != this.Head && curr.Prev.Freq <= node.Freq {
// 		curr = curr.Prev
// 	}
// 	node.Prev = curr.Prev
// 	curr.Prev.Next = node
// 	curr.Prev = node
// 	node.Next = curr
// }

// func (this *LFUCache) Increase(node *Node) {
//     newFreq := node.Freq + 1

//     node.Prev.Next = node.Next
//     node.Next.Prev = node.Prev

//     curr := node.Next

//     for curr.Prev != this.Head && curr.Prev.Freq <= newFreq {
//         curr = curr.Prev
//     }

//     node.Freq = newFreq

//     node.Prev = curr.Prev
//     node.Next = curr
//     curr.Prev.Next = node
//     curr.Prev = node
// }

// func (this *LFUCache) Get(key int) int {
// 	node, ok := this.Cache[key]
// 	if !ok {
// 		return -1
// 	}
// 	this.Increase(node)
// 	return node.Val
// }

// func (this *LFUCache) Put(key int, value int)  {
// 	if node, ok := this.Cache[key]; ok {
// 		node.Val = value
// 		this.Increase(node)
// 		return
// 	}
// 	node := &Node{
// 		Key: key,
// 		Val: value,
// 		Freq: 1,
// 	}
// 	if len(this.Cache) >= this.Cap {
// 		this.Remove()
// 	}
// 	this.Cache[key] = node
// 	this.Insert(node)
// }

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */
