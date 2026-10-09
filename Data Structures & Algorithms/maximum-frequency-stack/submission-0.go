type FreqStack struct {
	Freq map[int]int
	Stack [][]int
	MaxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		Freq: make(map[int]int),
		Stack: make([][]int, 1),
	}
}

func (this *FreqStack) Push(val int) {
	this.Freq[val]++
	freq := this.Freq[val]
	if freq == len(this.Stack) {
		this.Stack = append(this.Stack, []int{})
	}
	this.Stack[freq] = append(this.Stack[freq], val)

	this.MaxFreq = max(this.MaxFreq, freq)
}

func (this *FreqStack) Pop() int {
	stack := this.Stack[this.MaxFreq]
	val := stack[len(stack) - 1]
	this.Stack[this.MaxFreq] = stack[:len(stack) - 1]
	
	this.Freq[val]--

	if len(this.Stack[this.MaxFreq]) == 0 {
		this.MaxFreq--
	}

	return val 
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 