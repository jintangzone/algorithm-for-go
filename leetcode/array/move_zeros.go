package array

// 覆盖法
//func moveZeros(nums []int) {
//	nonZeroAt := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] != 0 {
//			nums[nonZeroAt] = nums[i]
//			if i != nonZeroAt {
//				nums[i] = 0
//			}
//			nonZeroAt++
//		}
//	}
//}

func moveZeros(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
}
