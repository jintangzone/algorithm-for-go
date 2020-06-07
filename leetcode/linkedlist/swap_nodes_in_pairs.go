package linkedlist

import "algorithm-for-go/base/structure/linkedlist"

func SwapNode(link *linkedlist.LinkedList) {
	if link.Head == nil || link.Head.Next == nil {
		return
	}
	pre := new(linkedlist.LinkNode)
	pre.Next = link.Head
	link.Head = link.Head.Next
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
	}
}
