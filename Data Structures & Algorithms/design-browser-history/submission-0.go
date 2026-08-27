type Node struct {
	Val string
	Next, Prev *Node
}

type BrowserHistory struct {
	Curr *Node
}

func Constructor(homepage string) BrowserHistory {
	node := &Node{
		Val: homepage,
	}
    return BrowserHistory{
		Curr: node,
	}
}

func (this *BrowserHistory) Visit(url string)  {
	this.Curr.Next = &Node{
		Val: url,
		Prev: this.Curr,
	}
	this.Curr = this.Curr.Next
}

func (this *BrowserHistory) Back(steps int) string {
	for this.Curr.Prev != nil && steps > 0 {
		this.Curr = this.Curr.Prev
		steps--
	}
	return this.Curr.Val
}

func (this *BrowserHistory) Forward(steps int) string {
	for this.Curr.Next != nil && steps > 0 {
		this.Curr = this.Curr.Next
		steps--
	}
	return this.Curr.Val
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */