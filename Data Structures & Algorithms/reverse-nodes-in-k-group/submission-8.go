/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	dummy := &ListNode{Next: head}
// 	groupPrev := dummy
// 	for {
// 		kthNode := groupPrev.getKth(k)
// 		if kthNode == nil {
// 			break
// 		}
// 		groupHead := kthNode.Next
// 		oldHead := groupPrev.Next 
// 		prev := kthNode.Next
// 		curr := groupPrev.Next
// 		for curr != groupHead {
// 			next := curr.Next
// 			curr.Next = prev
// 			prev = curr
// 			curr = next
// 		} 
// 		groupPrev.Next = prev
// 		groupPrev = oldHead
// 	}
// 	return dummy.Next
// }

// func(l *ListNode) getKth(k int) *ListNode {
// 	for l != nil && k > 0 { 
//		l = l.Next
// 		k--		
// 	}
// 	return l 
// }

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}
	newHead := reverseKGroup(tail, k)

	target := tail 
	prev := newHead
	curr := head
	for curr != target {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}