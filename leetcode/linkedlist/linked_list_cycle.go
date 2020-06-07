package linkedlist

import "algorithm-for-go/base/structure/linkedlist"

func HasCycle(link *linkedlist.LinkedList) bool {
	fast := link.Head
	slow := fast
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
