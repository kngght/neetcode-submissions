/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// func insertionSortList(head *ListNode) *ListNode {
//         // Фиктивный узел -- начало упорядоченного списка.
//     dummy := &ListNode{Next: head}
//     for head != nil {
//         curr := head 
//         // Сдвиг начала основного списка.
//         head = head.Next
//         // Отделение начального узла от основного списка.
//         curr.Next = nil
//         // Поиск узла, значение которого больше 
//         // значения текущего в упорядоченном списке.
//         // Если такого узла нет, текущий встанет на место последнего.
//         prev := dummy
//         for prev.Next != nil && curr.Val >= prev.Next.Val {
//             prev = prev.Next
//         }
//         // Вставка текущего узла в упорядоченный список. 
//         curr.Next = prev.Next
//         prev.Next = curr
//     }
//     return dummy.Next
// }

func insertionSortList(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    // Два указателя
    prev, curr := head, head.Next
    for curr != nil {
        if prev.Val <= curr.Val {
            prev, curr = curr, curr.Next
            continue
        }
        // Поиск нужной позиции с начала списка 
        tmp := dummy
        for tmp.Next.Val <= curr.Val && tmp.Next != nil {
            tmp = tmp.Next
        }
        // Смена позиций узлов
        prev.Next = curr.Next
        curr.Next = tmp.Next
        tmp.Next = curr
        curr = prev.Next
    }
    return dummy.Next
}
