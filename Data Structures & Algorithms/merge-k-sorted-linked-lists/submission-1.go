/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		c := len(lists)/2 + len(lists)%2
		merged := make([]*ListNode, 0, c)
		for i := 0; i < len(lists); i+=2 {
			if i+1 < len(lists) {
				merged = append(merged, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				merged = append(merged, lists[i])
			}
		}
		lists = merged
	}
	return lists[0]
}
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 == nil {
		head.Next = list2
	} else {
		head.Next = list1
	}
	return dummy.Next
}