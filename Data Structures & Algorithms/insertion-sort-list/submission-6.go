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

// func insertionSortList(head *ListNode) *ListNode {
//     dummy := &ListNode{}
//     for head != nil {
//         curr := head
//         head = head.Next
//         curr.Next = nil
//         prev := dummy
//         for prev.Next != nil && prev.Next.Val < curr.Val {
//             prev = prev.Next
//         }
//         curr.Next = prev.Next
//         prev.Next = curr
//     }
//     return dummy.Next
// }

// func insertionSortList(head *ListNode) *ListNode {
//     for curr := head; curr != nil; curr = curr.Next {
//         smallest := curr 
//         for tmp := curr; tmp != nil; tmp = tmp.Next {
//             if tmp.Val < smallest.Val {
//                 smallest = tmp
//             }
//         }
//         curr.Val, smallest.Val = smallest.Val, curr.Val
//     } 
//     return head
// }

func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{Next: head}
    prev, curr := head, head.Next
    for curr != nil {
        for curr.Val >= prev.Val {
            prev = curr
            curr = curr.Next
            if curr == nil {
                return dummy.Next
            }
        }
        tmp := dummy
        for tmp.Next.Val < curr.Val {
            tmp = tmp.Next
        }
        prev.Next = curr.Next
        curr.Next = tmp.Next
        tmp.Next = curr
        curr = prev.Next
    }
    return dummy.Next
}