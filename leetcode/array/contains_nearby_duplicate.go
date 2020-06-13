package array

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]bool)
	l := 0
	for r := 0; r < len(nums); r++ {
		if _, ok := set[nums[r]]; ok {
			return true
		}
		set[nums[r]] = true
		if r-l == k {
			delete(set, nums[l])
			l++
		}
	}
	return false
}
