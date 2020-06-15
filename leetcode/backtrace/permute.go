package backtrace

type Permute struct {
	res  [][]int
	used map[int]bool
}

func NewPermute() *Permute {
	return &Permute{
		res:  make([][]int, 0),
		used: make(map[int]bool),
	}
}

func (p *Permute) generatePermutation(nums []int, index int, sub []int) {
	if index == len(nums) {
		var subRes = make([]int, len(nums))
		copy(subRes, sub)
		p.res = append(p.res, subRes)
		return
	}
	for i := 0; i < len(nums); i++ {
		ok := p.used[i]
		if !ok {
			sub = append(sub, nums[i])
			p.used[i] = true
			p.generatePermutation(nums, index+1, sub)
			sub = sub[:len(sub)-1]
			delete(p.used, i)
		}
	}
	return
}

func permute(nums []int) [][]int {
	pm := NewPermute()
	if len(nums) == 0 {
		return pm.res
	}
	sub := make([]int, 0)
	pm.generatePermutation(nums, 0, sub)
	return pm.res
}
