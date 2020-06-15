package array

var num int
var memo []int

func fib1(n int) int {
	num++
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] == -1 {
		memo[n] = fib1(n-1) + fib1(n-2)
	}
	return memo[n]
}

func fib2(n int) int {
	mm := make([]int, n+1)
	mm[0] = 0
	mm[1] = 1
	for i := 2; i <= n; i++ {
		mm[i] = mm[i-1] + mm[i-2]
	}
	return mm[n]
}
