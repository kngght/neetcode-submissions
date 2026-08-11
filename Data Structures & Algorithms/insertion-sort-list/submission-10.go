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

// func insertionSortList(head *ListNode) *ListNode {
//     dummy := &ListNode{Next: head}
//     // Два указателя
//     prev, curr := head, head.Next
//     for curr != nil {
//         if prev.Val <= curr.Val {
//             prev, curr = curr, curr.Next
//             continue
//         }
//         // Поиск нужной позиции с начала списка 
//         tmp := dummy
//         for tmp.Next.Val <= curr.Val && tmp.Next != nil {
//             tmp = tmp.Next
//         }
//         // Смена позиций узлов
//         prev.Next = curr.Next
//         curr.Next = tmp.Next
//         tmp.Next = curr
//         curr = prev.Next
//     }
//     return dummy.Next
// }

// func insertionSortList(head *ListNode) *ListNode {
//     // Необходим для проверки первого элемента списка.
//     firstDummy := &ListNode{Next: head}
//     // Фиктивный узел для нового упорядоченного списка.
//     secondDummy := &ListNode{}
//     tail := secondDummy 
//     for firstDummy.Next != nil {
//         // Предыдущее значение необходимо для последующего узла.
//         prevMin, minNode := firstDummy, firstDummy.Next
//         // Данные указатели нужны для обхода основного списка
//         // и поиска значений для prevMin и minNode.
//         prev, curr := firstDummy, firstDummy.Next 
//         for curr != nil {
//             if curr.Val < minNode.Val {
//                 prevMin, minNode = prev, curr
//             }
//             prev, curr = curr, curr.Next
//         }
//         // Разрыв minNode от основного списка и добавление в упорядоченный.
//         prevMin.Next = minNode.Next
//         tail.Next = minNode
//         tail = tail.Next
//     }
//     // В добавленном узле может сохраняться ссылка на not nil узел.
//     tail.Next = nil
//     return secondDummy.Next
// }

func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    length := 0
    for i := head; i != nil; i = i.Next {
        length++
    }
    dummy := &ListNode{Next: head}
    for i := 1; i < length; i*=2 {
        prev := dummy
        curr := dummy.Next
        for curr != nil {
            left := curr
            right := split(left, i)
            curr = split(right, i)
            mergedHead, mergedTail := merge(left, right)
            prev.Next = mergedHead
            prev = mergedTail
        }
    } 
    return dummy.Next
}

func merge(list1, list2 *ListNode) (*ListNode, *ListNode) {
    dummy := &ListNode{}
    curr := dummy
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            curr.Next = list1
            list1 = list1.Next
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next
    }
    if list1 == nil {
        curr.Next = list2
    } else {
        curr.Next = list1
    }
    for curr.Next != nil {
        curr = curr.Next
    }
    return dummy.Next, curr
}

func split(head *ListNode, k int) *ListNode {
    for head != nil &&  k > 1 {
        head = head.Next
        k--
    } 
    if head == nil {
        return nil 
    }
    curr := head.Next
    head.Next = nil
    return curr
}