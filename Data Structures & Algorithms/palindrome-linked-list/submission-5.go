/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev 
		prev = slow 
		slow = next	
	}

	left := head
	right := prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}	
		left = left.Next	
		right = right.Next
	} 
	return true
}

// func isPalindrome(head *ListNode) bool {
// 	result := true
// 	var recurse func(l *ListNode)
// 	recurse = func(l *ListNode) {
// 		if l == nil {
// 			return
// 		}
// 		recurse(l.Next)
// 		if head.Val != l.Val {
// 			result = false
// 		}
// 		head = head.Next
// 	}
// 	recurse(head)
// 	return result
// }