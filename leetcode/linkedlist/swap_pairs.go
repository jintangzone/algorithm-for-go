package linkedlist

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	p := dummyHead
	for p.Next != nil && p.Next.Next != nil {
		a := p.Next
		b := a.Next
		p.Next, a.Next, b.Next = b, b.Next, a
		p = a
	}
	return dummyHead.Next
}
