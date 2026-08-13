/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// func rotateRight(head *ListNode, k int) *ListNode {
// 	if head == nil || head.Next == nil || k == 0 {
// 		return head
// 	}
// 	length := 1
// 	tail := head
// 	for tail.Next != nil {
// 		tail = tail.Next
// 		length++
// 	} 
// 	k %= length
// 	if k == 0 {
// 		return head
// 	}
// 	slow, fast := head, head
// 	for i := 0; i < k; i++ {
// 		fast = fast.Next
// 	}
// 	for fast.Next != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}
// 	fast.Next = head
// 	head = slow.Next
// 	slow.Next = nil
// 	return head
// }

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }
    tail := head  
    length := 1
    for tail.Next != nil {
        tail = tail.Next
        length++
    }
    k %= length
    if k == 0 {
        return head
    }
    tail.Next = head
    prevHead := head
    for i := 1; i < length-k; i++ {
        prevHead = prevHead.Next
    }
    head = prevHead.Next
    prevHead.Next = nil
    return head
}