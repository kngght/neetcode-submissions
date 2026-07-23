/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// length := 0
	// for curr := head; curr != nil; curr = curr.Next {
		// length++
	// }
// 
	// dummy := &ListNode{Next: head}
	// prev := dummy
	// for i := 1; i <= length-n; i++ {
		// prev = prev.Next
	// }
// 
	// prev.Next = prev.Next.Next
	// return dummy.Next 
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy 

	for i := 0; i < n; i++ {
		fast = fast.Next
	}	

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next 
}