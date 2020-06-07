package linkedlist

import "fmt"

type LinkNode struct {
	Value int
	Next  *LinkNode
}

type LinkedList struct {
	Head *LinkNode
}

func (ll *LinkedList) Print() {
	p := ll.Head
	for p != nil {
		fmt.Print(p.Value, "->")
		p = p.Next
	}
	fmt.Println("nil")
}

func NewLinkedList(arr []int) *LinkedList {
	if len(arr) <= 0 {
		return nil
	}
	p := &LinkNode{
		Value: arr[0],
		Next:  nil,
	}
	ll := &LinkedList{Head: p}
	for i := 1; i < len(arr); i++ {
		p.Next = &LinkNode{Value: arr[i], Next: nil}
		p = p.Next
	}
	return ll
}
