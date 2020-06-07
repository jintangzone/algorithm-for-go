package test

import (
	"algorithm-for-go/base/structure/stack"
	"testing"
)

func TestLinkedStack(t *testing.T) {
	ls := new(stack.LinkedStack)
	ls.Push(1)
	ls.Push(3)
	ls.Push(5)
	ls.Push(7)
	ls.Print()
	t.Log("peek:", ls.Peek())
	t.Log("pop:", ls.Pop())
	ls.Print()
	ls.Push(6)
	ls.Print()
}
