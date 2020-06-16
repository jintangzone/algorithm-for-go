package greedy

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	si, gi, res := len(s)-1, len(g)-1, 0
	for gi >= 0 && si >= 0 {
		if s[si] >= g[gi] {
			res++
			si--
			gi--
		} else {
			gi--
		}
	}
	return res
}
