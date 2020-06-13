package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	p := dummyHead
	q := p
	for p.Next != nil {
		if n <= 0 {
			q = q.Next
		}
		n--
		p = p.Next
	}
	q.Next = q.Next.Next
	return dummyHead.Next
}
