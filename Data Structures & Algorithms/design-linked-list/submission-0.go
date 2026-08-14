type MyLinkedList struct {
    head *ListNode
    tail *ListNode
    size int
}

type ListNode struct {
    val  int
    prev *ListNode
    next *ListNode
}

func Constructor() MyLinkedList {
    return MyLinkedList{
        head: nil,
        tail: nil,
        size: 0,
    }
}

// Получить значение узла по индексу (0-based)
func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    cur := this.head
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    return cur.val
}

// Добавить в начало
func (this *MyLinkedList) AddAtHead(val int) {
    newNode := &ListNode{val: val}
    if this.size == 0 {
        this.head = newNode
        this.tail = newNode
    } else {
        newNode.next = this.head
        this.head.prev = newNode
        this.head = newNode
    }
    this.size++
}

// Добавить в конец
func (this *MyLinkedList) AddAtTail(val int) {
    newNode := &ListNode{val: val}
    if this.size == 0 {
        this.head = newNode
        this.tail = newNode
    } else {
        this.tail.next = newNode
        newNode.prev = this.tail
        this.tail = newNode
    }
    this.size++
}

// Добавить по индексу (если index == size — в конец)
func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    if index == this.size {
        this.AddAtTail(val)
        return
    }

    // Ищем узел, который будет стоять после нового (т.е. текущий узел на позиции index)
    cur := this.head
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    // Вставляем перед cur
    prevNode := cur.prev
    newNode := &ListNode{val: val, prev: prevNode, next: cur}
    prevNode.next = newNode
    cur.prev = newNode
    this.size++
}

// Удалить по индексу
func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    if this.size == 1 {
        this.head = nil
        this.tail = nil
        this.size = 0
        return
    }

    // Удаляем голову
    if index == 0 {
        this.head = this.head.next
        this.head.prev = nil
        this.size--
        return
    }

    // Удаляем хвост
    if index == this.size-1 {
        this.tail = this.tail.prev
        this.tail.next = nil
        this.size--
        return
    }

    // Удаляем из середины
    cur := this.head
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    cur.prev.next = cur.next
    cur.next.prev = cur.prev
    this.size--
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