type MyHashSet struct {
	set []bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		set: make([]bool, 1_000_001),
	}	
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.set[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.set[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 