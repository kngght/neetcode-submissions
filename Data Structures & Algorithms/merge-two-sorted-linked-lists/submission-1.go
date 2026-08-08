/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	dummy := &ListNode{}
	head := dummy
	first := list1
	second := list2
	for {
		if second == nil {
			head.Next = first 
			break
		}
		if first == nil {
			head.Next = second
			break
		}
		if first.Val < second.Val {
			head.Next = first
			head = head.Next
			first = first.Next
		} else {
			head.Next = second 
			head = head.Next
			second = second.Next
		}
	}
	return dummy.Next
}

