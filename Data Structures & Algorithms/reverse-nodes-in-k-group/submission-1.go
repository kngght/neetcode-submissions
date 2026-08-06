/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroup := dummy
	for {
		kNode := prevGroup
		for i := 0; i < k; i++ {
			if kNode == nil {
				break
			}
			kNode = kNode.Next
		}
		if kNode == nil {
			break
		}
		oldHead := prevGroup.Next	
		target := kNode.Next 
		prev := kNode.Next
		curr :=	prevGroup.Next 
		for curr != target {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		prevGroup.Next = prev
		prevGroup = oldHead
	}
	return dummy.Next
}