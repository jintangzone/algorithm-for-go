package array

func sortColors(nums []int) {
	n := len(nums)
	zero := -1 // nums[0...zero] == 0，初始zero=-1，故意使其无效
	two := n   // nums[two...n-1] == 2，初始two=n，故意使其无效
	i := 0
	for i < two {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[two-1], nums[i] = nums[i], nums[two-1]
			two--
		} else { // nums[i] == 0
			nums[zero+1], nums[i] = nums[i], nums[zero+1]
			i++
			zero++
		}
	}

}
