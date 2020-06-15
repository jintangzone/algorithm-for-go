package backtrace

type Combine struct {
	res [][]int
}

func (c *Combine) generateCombinations(n, k, start int, sub []int) {
	if len(sub) == k {
		subRes := make([]int, k)
		copy(subRes, sub)
		c.res = append(c.res, subRes)
		return
	}
	for i := start; i <= n-(k-len(sub))+1; i++ {
		sub = append(sub, i)
		c.generateCombinations(n, k, i+1, sub)
		sub = sub[:len(sub)-1]
	}
	return
}

func NewCombine() *Combine {
	return &Combine{res: make([][]int, 0)}
}

func combine(n int, k int) [][]int {
	c := NewCombine()
	if n <= 0 || k <= 0 || k > n {
		return c.res
	}
	sub := make([]int, 0)
	c.generateCombinations(n, k, 1, sub)
	return c.res
}
