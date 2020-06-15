package dynamic

/**
 * 状态定义: 考虑将n个物品填满容量为C的背包
 * 状态转移：F(i, c) = F(i-1,c)||F(i-1, c-w(i))
 */

/**
 * 递归方式
 */
//func canPartition(nums []int) bool {
//	sum := 0
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//	}
//
//	if sum%2 != 0 {
//		return false
//	}
//
//	return tryPartition(nums, len(nums)-1, sum/2)
//}
//
//// 使用nums[0...index]，是否可以完全填充一个容量为sum的背包
//func tryPartition(nums []int, index int, sum int) bool {
//	if sum == 0 {
//		return true
//	}
//	if sum < 0 || index < 0 {
//		return false
//	}
//	return tryPartition(nums, index-1, sum) ||
//		tryPartition(nums, index-1, sum-nums[index])
//}

/**
 * 记忆化搜索
 */
//var memo [][]int
//
//func canPartition(nums []int) bool {
//	sum := 0
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//	}
//	if sum%2 != 0 {
//		return false
//	}
//	memo = make([][]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		row := make([]int, sum/2+1)
//		for j := 0; j < sum/2+1; j++ {
//			row[j] = -1
//		}
//		memo[i] = row
//	}
//	return tryPartition(nums, len(nums)-1, sum/2)
//}
//
//// 使用nums[0...index]，是否可以完全填充一个容量为sum的背包
//func tryPartition(nums []int, index int, sum int) bool {
//	if sum == 0 {
//		return true
//	}
//	if sum < 0 || index < 0 {
//		return false
//	}
//	if memo[index][sum] != -1 {
//		return memo[index][sum] == 1
//	}
//	if tryPartition(nums, index-1, sum) ||
//		tryPartition(nums, index-1, sum-nums[index]) {
//		memo[index][sum] = 1
//	} else {
//		memo[index][sum] = 0
//	}
//	return memo[index][sum] == 1
//}

/**
 * 动态规划方式
 */

// dp[j] = dp[j] || dp[j - nums[i]];
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	n := len(nums)
	C := sum / 2
	dp := make([]bool, C+1)
	for i := 0; i <= C; i++ {
		dp[i] = nums[0] == i
	}
	// dp[j] = dp[j] || dp[j-nums[i]]
	for i := 1; i < n; i++ {
		for j := C; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[C]
}
