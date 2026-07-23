/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num int
	dummy := &ListNode{}
	newList := dummy
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + num
		num = val / 10
		newList.Next = &ListNode{Val: val % 10}
		newList = newList.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + num
		num = val / 10
		newList.Next = &ListNode{Val: val % 10}
		newList = newList.Next
		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + num
		num = val / 10
		newList.Next = &ListNode{Val: val % 10}
		newList = newList.Next
		l2 = l2.Next
	}
	if num > 0 {
		newList.Next = &ListNode{Val: num}
	}
	return dummy.Next
}
