package stack

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedStack struct {
	Head   *Node
	length int
}

func (ls *LinkedStack) Peek() interface{} {
	return ls.Head.Value
}

func (ls *LinkedStack) Push(value interface{}) {
	node := &Node{Value: value, Next: ls.Head}
	ls.Head = node
	ls.length++
}

func (ls *LinkedStack) Pop() interface{} {
	if ls.length == 0 {
		return nil
	}
	ret := ls.Head
	ls.Head = ls.Head.Next
	ret.Next = nil
	ls.length--
	return ret.Value
}

func (ls *LinkedStack) Print() {
	for p := ls.Head; p != nil; p = p.Next {
		fmt.Print(p.Value)
		if p.Next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
