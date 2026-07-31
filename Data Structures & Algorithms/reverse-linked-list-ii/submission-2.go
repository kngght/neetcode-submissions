/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	var rightHead *ListNode
// 	var reverse func(head *ListNode, n int) *ListNode
// 	reverse = func(head *ListNode, n int) *ListNode {
// 		if n == 1 {
// 			rightHead = head.Next
// 			return head
// 		}
// 		newHead := reverse(head.Next, n-1)	
// 		head.Next.Next = head
// 		head.Next = rightHead
// 		return newHead
// 	}
// 	if left == 1 {
// 		return reverse(head, right)
// 	}
// 	head.Next = reverseBetween(head.Next, left-1, right-1)
// 	return head
// }


func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prevSublist := dummy
	for i := 1; i < left; i++ {
		prevSublist = prevSublist.Next
	}
	var prev *ListNode
	curr := prevSublist.Next
	for i := 0; i <= right - left; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	prevSublist.Next.Next = curr
	prevSublist.Next = prev
	return dummy.Next
}