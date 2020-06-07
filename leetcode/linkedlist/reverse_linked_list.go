package linkedlist

import "algorithm-for-go/base/structure/linkedlist"

func ReverseLinkedList(link *linkedlist.LinkedList) {
	var prev *linkedlist.LinkNode
	cur := link.Head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	link.Head = prev
}
