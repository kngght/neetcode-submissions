/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    groupPrev := dummy
    for {
        kthNode := getKth(groupPrev, k)
        if kthNode == nil {
            break
        }
        oldHead := groupPrev.Next
        groupPrev.Next = reverse(groupPrev.Next, kthNode.Next) 
        groupPrev = oldHead
    }
    return dummy.Next
}
func getKth(head *ListNode, k int) *ListNode {
    for head != nil && k > 0 {
        head = head.Next
        k--
    }
    return head
}
func reverse(curr, prev *ListNode) *ListNode {
    target := prev
    for curr != target {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	nextHead := head
// 	for i := 0; i < k; i++ {
// 		if nextHead == nil {
// 			return head
// 		}
// 		nextHead = nextHead.Next
// 	}
// 	prev := reverseKGroup(nextHead, k) 
// 	curr := head
// 	for curr != nextHead {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	return prev
// }