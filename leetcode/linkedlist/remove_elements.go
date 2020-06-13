package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	p := dummyHead
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return dummyHead.Next
}
