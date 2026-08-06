/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy	
	
	for {
		kNode := prev.get(k) 
		if kNode == nil {
			break
		}
		groupHead := prev.Next
		prev.Next = kNode.reverse(groupHead)
		prev = groupHead
	}
	return dummy.Next
}

func (l *ListNode) get(k int) *ListNode {
	for l != nil && k > 0 {
		l = l.Next
		k--
	}
	return l
}

func (l *ListNode) reverse(groupHead *ListNode) *ListNode {
	target := l.Next
	prev := l.Next
	curr := groupHead
	for curr != target {
		next := curr.Next
		curr.Next = prev 
		prev = curr
		curr = next	
	}
	return prev
}