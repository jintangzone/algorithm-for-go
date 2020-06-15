package dynamic

/**
 * 状态定义: 将n进行分割（至少分割成2部分），得到最大的乘机
 * 状态转移：
		f(n) = max{
			1*f(n-1),
			2*f(n-2),
			...,
			i*f(n-i)
		}
*/

/**
 * 递归方式
 */
//func integerBreak(n int) int {
//	return breakInteger(n)
//}
//
//// 将n进行分割（至少分割两部分），可以获得最大成绩
//func breakInteger(n int) int {
//	if n == 1 {
//		return 1
//	}
//	res := -1
//	for i := 1; i <= n-1; i++ {
//		// i + (n-i)
//		res = max(res, max(i*(n-i), i*breakInteger(n-i)))
//	}
//	return res
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

/**
 * 记忆化搜索方式
 */
//var memo []int
//
//func integerBreak(n int) int {
//	memo = make([]int, n+1)
//	for i := 0; i <= n; i++ {
//		memo[i] = -1
//	}
//	return breakInteger(n)
//}
//
//// 将n进行分割（至少分割两部分），可以获得最大成绩
//func breakInteger(n int) int {
//	if n == 1 {
//		return 1
//	}
//	if memo[n] != -1 {
//		return memo[n]
//	}
//
//	res := -1
//	for i := 1; i <= n-1; i++ {
//		// i + (n-i)
//		res = max(res, max(i*(n-i), i*breakInteger(n-i)))
//		memo[n] = res
//	}
//	return res
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

/**
 * 动态规划方式
 */
func integerBreak(n int) int {
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}
	memo[1] = 1
	for i := 2; i <= n; i++ {
		// 求解memo[i]
		for j := 1; j <= i-1; j++ {
			// j + (i-j)
			memo[i] = max(memo[i], max(j*(i-j), j*memo[i-j]))
		}
	}
	return memo[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
