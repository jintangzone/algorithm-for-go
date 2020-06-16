package greedy

import (
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	t.Log(findContentChildren(g, s))
}
