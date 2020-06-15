package dynamic

/**
 * 状态定义：考虑偷取[x...n-1]范围里的房子（函数定义）
 * 状态转移：
	dp(i) = max{dp
	f(0) = max{
		v(0)+f(2, n-1),
		v(1)+f(3, n-1),
		v(2)+f(4, n-1),
		...,
		v(n-
		...,
		v(n-3)+f(n-1, n-1),
		v(n-2),
		v(n-1)
	}
*/

/**
 * 递归方式
 */
// 考虑偷取[x...n-1]范围里的房子（函数定义）
//func rob(nums []int) int {
//	return tryRob(nums, 0)
//}
//
//// 考虑偷取[index...len(nums)]这个范围的所有房子
//func tryRob(nums []int, index int) int {
//	if index >= len(nums) {
//		return 0
//	}
//	res := 0
//	// f(index, n-1) = max{v(0) + f(2, n-1), ... v(index)+f(index+2, n-1),...}
//	for i := index; i < len(nums); i++ {
//		res = maxV(res, nums[i]+tryRob(nums, i+2))
//	}
//	return res
//}
//
//func maxV(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

/**
 * 记忆化搜索
 */

//var memo []int
//
//func rob(nums []int) int {
//	memo = make([]int, len(nums)+1)
//	for i := 0; i <= len(nums); i++ {
//		memo[i] = -1
//	}
//	return tryRob(nums, 0)
//}
//
//// 考虑偷取[index...len(nums)]这个范围的所有房子
//func tryRob(nums []int, index int) int {
//	if index >= len(nums) {
//		return 0
//	}
//
//	if memo[index] != -1 {
//		return memo[index]
//	}
//	res := 0
//	// f(index, n-1) = max{v(0) + f(2, n-1), ... v(index)+f(index+2, n-1),...}
//	for i := index; i < len(nums); i++ {
//		res = maxV(res, nums[i]+tryRob(nums, i+2))
//	}
//	memo[index] = res
//	return res
//}
//
//func maxV(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

/**
 * 动态规划
 */

/**
 * 状态定义：考虑偷取[x...n-1]范围里的房子（函数定义）
 * 状态转移：
	dp(i) = max{dp
	f(0) = max{
		v(0)+f(2, n-1),
		v(1)+f(3, n-1),
		v(2)+f(4, n-1),
		...,
		v(n-
		...,
		v(n-3)+f(n-1, n-1),
		v(n-2),
		v(n-1)
	}
*/
//func rob(nums []int) int {
//	n := len(nums)
//	if n == 0 {
//		return 0
//	}
//	// memo[i] 表示考虑抢劫nums[i...n-1]所获得的最大收益
//	memo := make([]int, n)
//	for i := 0; i < n; i++ {
//		memo[i] = -1
//	}
//	memo[n-1] = nums[n-1]
//	for i := n - 2; i >= 0; i-- {
//		// memo[i]
//		for j := i; j < n; j++ {
//			tmp := 0
//			if j+2 < n {
//				tmp = memo[j+2]
//			}
//			memo[i] = maxR(memo[i], nums[j]+tmp)
//		}
//	}
//	return memo[0]
//}
//
//func maxR(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

/**
 * 状态定义：考虑偷取[0...i]范围内房子，获得最大收益
 * 状态转移：
 * 		dp(i) = max{dp(i-2)+nums[i], dp(i-1)}
 */
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	//dp := make([]int, n)
	//dp[0] = nums[0]
	//dp[1] = maxR(nums[0],nums[1])
	//for i := 2; i < n; i++ {
	//	dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	//}
	//return dp[n-1]
	first := nums[0]
	second := maxR(nums[0], nums[1])
	for i := 2; i < n; i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func maxR(a, b int) int {
	if a > b {
		return a
	}
	return b
}
