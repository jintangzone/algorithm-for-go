package test

import (
	"algorithm-for-go/leetcode/stackandqueue"
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "(([[]])"
	fmt.Println(s, " valid is ", stackandqueue.IsValid(s))
}
