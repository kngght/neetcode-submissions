type MyLinkedList struct {
    Head *Node 
    Tail *Node
    Size int
}

type Node struct {
    Val int
    Next *Node
}

func Constructor() MyLinkedList {
    return MyLinkedList{
        Head: nil,
        Tail: nil,
        Size: 0,
    } 
}


func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.Size {
        return -1
    } 
    curr := this.Head
    for i := 0; i < index; i++ {
        curr = curr.Next
    }
    return curr.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    if this.Size == 0 {
        node := &Node{Val: val}
        this.Head = node
        this.Tail = node
    } else {
        node := &Node{Val: val, Next: this.Head}
        this.Head = node
    }
    this.Size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    if this.Size == 0 {
        node := &Node{Val: val}
        this.Head = node
        this.Tail = node
    } else {
        node := &Node{Val: val}
        this.Tail.Next = node
        this.Tail = this.Tail.Next
    }
    this.Size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.Size {
        return 
    } 
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    if index == this.Size {
        this.AddAtTail(val)
        return
    }
    prev := this.Head
    for i := 0; i < index-1; i++ {
        prev = prev.Next
    }
    node := &Node{Val: val, Next: prev.Next}
    prev.Next = node
    this.Size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.Size {
        return 
    } 
    if this.Size == 1 {
        this.Head = nil
        this.Tail = nil
        this.Size = 0
        return
    }
    if index == 0 {
        this.Head = this.Head.Next
        if this.Size == 1 {
            this.Tail = nil
        }
        this.Size--
        return
    }
    prev := this.Head
    for i := 0; i < index-1; i++ {
        prev = prev.Next
    }
    if index == this.Size-1 {
        this.Tail = prev
    }
    prev.Next = prev.Next.Next
    this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */