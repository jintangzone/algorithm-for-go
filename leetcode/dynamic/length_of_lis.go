package dynamic

/**
 * 状态定义: LIS(i)表示以第i个数字为结尾的最长上升子序列的长度
			LIS(i)表示[0...i]的范围内，选择数字nums[i]可以获得最长上升子序列的长度
 * 状态转移: LIS(i) = max{1+LIS(j) if nums[i] > nums[j]}(i > j)
*/

/**
 * 动态规划
 */
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		res = max(res, dp[i])
	}
	return res
}
