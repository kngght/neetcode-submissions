/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prevLeft := dummy
	for i := 1; i < left; i++ {
		prevLeft = prevLeft.Next
	}
	leftHead := prevLeft.Next
	prevLeft.Next = nil

	prevRight := leftHead
	for i := 0; i < right-left; i++ {
		prevRight = prevRight.Next
	}
	rightHead := prevRight.Next
	prevRight.Next = nil

	var prev *ListNode
	curr := leftHead
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	prevLeft.Next = prev
	leftHead.Next = rightHead 

	return dummy.Next
}