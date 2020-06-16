package greedy

import (
	"sort"
)

//type intervals struct {
//	ii []*interval
//}
//
//func (in *intervals) Len() int {
//	return len(in.ii)
//}
//
//func (in *intervals) Swap(i, j int) {
//	in.ii[i], in.ii[j] = in.ii[j], in.ii[i]
//}
//
//func (in *intervals) Less(i, j int) bool {
//	a := in.ii[i]
//	b := in.ii[j]
//	if a.end != b.end {
//		return a.start < b.start
//	}
//	return a.end < b.end
//}
//
//func (in *intervals) Get(index int) *interval {
//	return in.ii[index]
//}
//
//func NewIntervals(intervalsGrid [][]int) *intervals {
//	in := new(intervals)
//	for i := 0; i < len(intervalsGrid); i++ {
//		in.ii = append(in.ii, NewInterval(intervalsGrid[i][0], intervalsGrid[i][1]))
//	}
//	return in
//}
//
//type interval struct {
//	start, end int
//}
//
//func NewInterval(start, end int) *interval {
//	return &interval{
//		start: start,
//		end:   end,
//	}
//}

/**
 * 动态规划
 */
//func eraseOverlapIntervals(intervals [][]int) int {
//	n := len(intervals)
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][1] < intervals[j][1]
//	})
//	dp := make([]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = 1
//	}
//	for i := 0; i < n; i++ {
//		for j := 0; j < i; j++ {
//			if intervals[i][0] >= intervals[j][1] {
//				dp[i] = max(dp[i], 1+dp[j])
//			}
//		}
//	}
//	res := 0
//	for i := 0; i < len(dp); i++ {
//		res = max(res, dp[i])
//	}
//	return n - res
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

/**
 * 贪心算法
 */
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res, pre := 1, 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[pre][1] {
			res++
			pre = i
		}
	}
	return len(intervals) - res
}
