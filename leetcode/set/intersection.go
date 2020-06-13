package set

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {

	set := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range nums1 {
		set[v] = true
	}

	for _, v := range nums2 {
		if set[v] == true {
			res = append(res, v)
			set[v] = false
		}
	}

	return res
}
