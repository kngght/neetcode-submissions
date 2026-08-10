/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// func insertionSortList(head *ListNode) *ListNode {
//     dummy := &ListNode{}
//     newList := dummy
//     for head != nil {
//         prevMin, minNode := (*ListNode)(nil), head
//         prev, curr := head, head.Next
//         for curr != nil {
//             if curr.Val < minNode.Val {
//                 prevMin = prev
//                 minNode = curr
//             }
//             prev = curr
//             curr = curr.Next
//         }
//         if prevMin == nil {
//             head = head.Next 
//         } else {
//             prevMin.Next = minNode.Next
//         }
//         newList.Next = minNode
//         newList = newList.Next
//     }
//     newList.Next = nil
//     return dummy.Next
// }

func insertionSortList(head *ListNode) *ListNode {
    dummy := &ListNode{}
    for head != nil {
        curr := head
        head = head.Next
        curr.Next = nil
        prev := dummy
        for prev.Next != nil && prev.Next.Val <= curr.Val {
            prev = prev.Next
        }
        curr.Next = prev.Next
        prev.Next = curr
    }
    return dummy.Next
}