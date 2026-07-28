/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	dummy := &ListNode{Next: head}
// 	slow, fast := dummy, dummy 

// 	for i := 0; i < n; i++ {
// 		fast = fast.Next
// 	}	

// 	for fast.Next != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}
// 	slow.Next = slow.Next.Next
// 	return dummy.Next 
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		 return nil
	}	
	dummy := &ListNode{Next: head}
	var recurse func(l *ListNode)
	recurse = func(l *ListNode) {
		if l.Next == nil {
			return
		}
		recurse(l.Next)
		n--
		if n == 0 {
			l.Next = l.Next.Next
		}
	}
	recurse(dummy)
	return dummy.Next 
}