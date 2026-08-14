/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
		dummy := &ListNode{}
	curr := dummy
	var carry int
	for l1 != nil && l2 != nil {
		sum := carry + l1.Val + l2.Val
		carry = sum/10
		curr.Next = &ListNode{Val: sum%10}
		curr = curr.Next
		l1, l2 = l2.Next, l1.Next
	}
	for l1 != nil {
		sum := carry + l1.Val
		carry = sum/10
		curr.Next = &ListNode{Val: sum%10}
		curr = curr.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := carry + l2.Val
		carry = sum/10
		curr.Next = &ListNode{Val: sum%10}
		curr = curr.Next
		l2 = l2.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	dummy.Next = reverse(dummy.Next)
	return dummy.Next	
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
