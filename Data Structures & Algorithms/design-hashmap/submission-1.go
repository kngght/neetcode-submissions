type Node struct {
    Key, Val int
    Next *Node
}

type LinkedList struct {
    Head *Node
}

func (l *LinkedList) Get(key int) int {
    for i := l.Head; i != nil; i = i.Next {
        if i.Key == key {
            return i.Val
        }
    }
    return -1
}

func (l *LinkedList) Put(key, val int) {
    for i := l.Head; i != nil; i = i.Next {
        if i.Key == key {
            i.Val = val
            return
        }
    }
    node := &Node{Key: key, Val: val, Next: l.Head}
    l.Head = node
}

func (l *LinkedList) Remove(key int) {
    if l.Head == nil {
        return
    } 

    if l.Head.Key == key {
        l.Head = l.Head.Next
        return
    }
    
    for i := l.Head; i.Next != nil; i = i.Next {
        if i.Next.Key == key {
            i.Next = i.Next.Next
            return
        }
    }
}

type MyHashMap struct {
    n int
    Buckets []*LinkedList
}


func Constructor() MyHashMap {
    n := 991
    buckets := make([]*LinkedList, n)
    for i := range buckets {
        buckets[i] = &LinkedList{}
    }
    return MyHashMap{
        n: n,
        Buckets: buckets,
    } 
}


func (this *MyHashMap) Put(key int, val int)  {
    hashKey := this.Hash(key)
    this.Buckets[hashKey].Put(key, val)
}


func (this *MyHashMap) Get(key int) int {
    hashKey := this.Hash(key)
    return this.Buckets[hashKey].Get(key)
}


func (this *MyHashMap) Remove(key int)  {
    hashKey := this.Hash(key)
    this.Buckets[hashKey].Remove(key)
}

func (this *MyHashMap) Hash(key int) int {
    return key%this.n
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */