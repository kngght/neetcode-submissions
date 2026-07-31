/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var rightHead *ListNode
	var reverse func(head *ListNode, n int) *ListNode
	reverse = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			rightHead = head.Next
			return head
		}
		newHead := reverse(head.Next, n-1)	
		head.Next.Next = head
		head.Next = rightHead
		return newHead
	}
	if left == 1 {
		return reverse(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}