package array

func minSubArrayLen(s int, nums []int) int {
	l := 0
	r := -1 // 定义nums[l..r]滑动窗口，为了初始化时区间[l...r]不包含元素，所以r取值-1
	sum := 0
	n := len(nums)
	res := n + 1

	for l < n {
		if r+1 < n && sum < s {
			sum += nums[r+1]
			r++
		} else {
			sum -= nums[l]
			l++
		}

		if sum >= s {
			if res > (r-l)+1 {
				res = (r - l) + 1
			}
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}
