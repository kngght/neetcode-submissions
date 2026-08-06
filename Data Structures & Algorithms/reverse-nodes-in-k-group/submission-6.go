/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := head
	prev := head
	curr := head
	flag := true

	for head != nil {
		for range k {
			if head == nil {
				return dummy
			}
			head = head.Next
		}

		reversed := reverse(curr, head, k) 
		if flag {
			dummy = reversed
			flag = false
		} else {
			prev.Next = reversed
		}
		prev = curr
		curr = head
	}
	return dummy
}

func reverse(head, tail *ListNode, k int) *ListNode {
	newHead := tail
	for range k {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}