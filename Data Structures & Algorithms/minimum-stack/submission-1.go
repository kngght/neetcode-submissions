type MinStack struct {
	Stack []int
	Min int
}

func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, 0)
		this.Min = val
	} else {
		this.Stack = append(this.Stack, val - this.Min)
		if val < this.Min {
			this.Min = val
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return 
	}

	pop := this.Stack[len(this.Stack) - 1]
	this.Stack = this.Stack[:len(this.Stack) - 1]

	if pop < 0 {
		this.Min -= pop
	}
}

func (this *MinStack) Top() int {
	top := this.Stack[len(this.Stack) - 1]
	if top > 0 {
		return top + this.Min
	}
	return this.Min
}

func (this *MinStack) GetMin() int {
	return this.Min
}
