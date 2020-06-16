package greedy

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	t.Log(eraseOverlapIntervals(intervals))
}
