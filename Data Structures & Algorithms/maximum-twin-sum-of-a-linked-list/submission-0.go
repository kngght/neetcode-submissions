/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    } 
    left, right := head, reverse(slow)
    sum := left.Val + right.Val
    for right != nil {
        if left.Val + right.Val > sum {
            sum = left.Val + right.Val
        }
        left = left.Next
        right = right.Next
    }
    return sum
}

func reverse(curr *ListNode) *ListNode {
    prev := (*ListNode)(nil)
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}