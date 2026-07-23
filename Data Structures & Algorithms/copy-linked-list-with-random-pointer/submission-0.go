/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	for i := head; i != nil; i = i.Next.Next {
		next := i.Next
		i.Next = &Node{Val: i.Val, Next: next}
	}

	for i := head; i != nil; i = i.Next.Next {
		if i.Random != nil {
			i.Next.Random = i.Random.Next
		}
	}

	dummy := &Node{}
	prev := dummy
	for i := head; i != nil; i = i.Next {
		next := i.Next
		i.Next = i.Next.Next
		prev.Next = next
		prev = next
	}
	return dummy.Next
}
