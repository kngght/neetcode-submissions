/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return 
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	second := slow.Next
	slow.Next = nil	
	second = reverse(second)
	first := head
	for second != nil {
		firstNext := first.Next
		secondNext := second.Next
		first.Next = second
		second.Next = firstNext
		first = firstNext
		second = secondNext
	}
}
 
func reverse(curr *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}	
	return prev
}