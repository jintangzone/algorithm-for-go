package test

import (
	"algorithm-for-go/base/structure/stack"
	"testing"
)

func TestArrayStack(t *testing.T) {
	as := stack.NewArrayStack()
	as.Push(1)
	as.Push(2)
	as.Push(3)
	t.Log(as.Arr)
	t.Log("peek:", as.Peek())
	t.Log("pop:", as.Pop())
	t.Log(as.Arr)
}
