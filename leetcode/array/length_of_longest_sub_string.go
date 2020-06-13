package array

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	l := 0
	r := -1 // 定义[l...r]为滑动窗口，r==-1故意为了让初始化区间[l...r]无效
	res := 0
	freq := make(map[uint8]int)
	for l < n {
		if r < n-1 && freq[s[r+1]] == 0 {
			freq[s[r+1]]++
			r++
		} else {
			freq[s[l]]--
			l++
		}
		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}
