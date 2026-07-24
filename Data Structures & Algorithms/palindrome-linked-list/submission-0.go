/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		hare = hare.Next.Next
		tortoise = tortoise.Next
	}
	
	var prev *ListNode 
	curr := tortoise
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	} 

	left, right := head, prev
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	return true
}
