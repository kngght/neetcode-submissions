type Pair struct {
	Price, Span int
}

type StockSpanner struct {
	Stack []Pair
}

func Constructor() StockSpanner {
	return StockSpanner{Stack: []Pair{}}
}

func (this *StockSpanner) Next(price int) int {
	span := 1	

	for len(this.Stack) > 0 && this.Stack[len(this.Stack) - 1].Price <= price {
		span += this.Stack[len(this.Stack) - 1].Span
		this.Stack = this.Stack[:len(this.Stack) - 1]
	}
	this.Stack = append(this.Stack, Pair{price, span})
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 