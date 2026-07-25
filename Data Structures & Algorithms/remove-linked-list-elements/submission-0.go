/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	left, right := dummy, dummy.Next

	for right != nil {
		if right.Val == val {
			left.Next = right.Next
			right = right.Next
		} else {
			left = right
			right = right.Next
		}
	}		
	return dummy.Next
}
