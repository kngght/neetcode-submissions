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
        tmp := dummy.Next 
        prev := dummy
        curr := head
        head = head.Next
        curr.Next = nil
        if tmp != nil {
            for {
            if curr.Val < tmp.Val {
                curr.Next = tmp 
                prev.Next = curr 
                break
            }
            if tmp.Next == nil {
                tmp.Next = curr
                break
            }
            prev = prev.Next
            tmp = tmp.Next
            }
        } else {
        curr.Next = tmp
        prev.Next = curr
        }
    }
    return dummy.Next
}