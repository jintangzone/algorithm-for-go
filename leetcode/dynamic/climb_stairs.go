package dynamic

/**
 * 递归方式
 */

//func climbStairs(n int) int {
//	if n == 1 {
//		return 1
//	}
//	if n == 2 {
//		return 2
//	}
//	return climbStairs(n-1) + climbStairs(n-2)
//}

/*
 * 记忆化搜索
 */
//var memo []int
//
//func climbStairs(n int) int {
//	memo = make([]int, n+1)
//	for i := 0; i <= n; i++ {
//		memo[i] = -1
//	}
//	return calcWays(n)
//}
//
//func calcWays(n int) int {
//	if n == 1 {
//		return 1
//	}
//	if n == 2 {
//		return 2
//	}
//	if memo[n] == -1 {
//		memo[n] = calcWays(n-1) + calcWays(n-2)
//	}
//	return memo[n]
//}

/**
 * 动态规划方式
 */

//func climbStairs(n int) int {
//	memo := make([]int, n+1)
//	memo[0] = 1
//	memo[1] = 1
//	for i := 2; i <= n; i++ {
//		memo[i] = memo[i-1] + memo[i-2]
//	}
//	return memo[n]
//}

/**
 * 滚动数组
 */

func climbStairs(n int) int {
	a, b, c := 0, 0, 1
	for i := 1; i <= n; i++ {
		a, b, c = b, c, a+b
	}
	return c
}
