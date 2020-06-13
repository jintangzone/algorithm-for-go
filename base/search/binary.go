package search

//func BinarySearch(arr []int, target int) int {
//
//	// 定义：在[l...r]的范围里寻找target
//	l := 0
//	r := len(arr) - 1
//
//	// 当l==r时，区间[l...r]依然是有效的, 所以是l <= r
//	for l <= r {
//		mid := l + (r-l)/2 // mid := (l+r)/2 的求法会有整型溢出的bug
//		if arr[mid] == target {
//			return mid
//		}
//		if target > arr[mid] {
//			l = mid + 1 // target在[mid+1...r]中
//		} else {
//			r = mid - 1 // target在[l...mid-1]中
//		}
//	}
//	return -1
//}
//
//func BinarySearch2(arr []int, target int) int {
//	// 定义：在[l...r)内查找target
//	l := 0
//	r := len(arr)
//
//	// 当l == r时，区间[l..r)无效，举例：[2,2)，l==2，l==r==2，与开区间矛盾
//	for l < r {
//		mid := l + (r-l)/2 // mid := (l+r)/2 的求法会有整型溢出的bug
//		if arr[mid] == target {
//			return mid
//		}
//		if target > arr[mid] {
//			// target在[mid+1...r)范围中
//			l = mid + 1
//		} else {
//			// target在[l...mid)范围中
//			r = mid
//		}
//	}
//	return -1
//}

func BinarySearch(nums []int, target int) int {

	// 定义在[l...r]范围内查找target

	l := 0
	r := len(nums) - 1

	// 当l==r时，区间[l...r]时有效的，如区间[2,2]，有一个元素就是2
	for l <= r {
		mid := l + (r-l)/2 // mid := (l+r)/2 会有大数相加的整型溢出
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			// 说明target在[mid+1...r]的范围内查找
			l = mid + 1
		} else {
			// 说明target在[l...mid-1]的范围内查找
			r = mid - 1
		}
	}
	return -1
}

func BinarySearch2(nums []int, target int) int {
	// 定义：在[l...r)范围内查找target
	l := 0
	r := len(nums)
	// 当l==r时，区间[l...r)不是一个有效区间，如区间[2,2)，表示没有任何元素，所以不能l==r
	for l < r {
		mid := l + (r-l)/2 // mid := (l+r)/2会有大数相加的整型溢出问题
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			// 说明target在[mid+1...r)的范围内查找
			l = mid + 1
		} else {
			// 说明target在[l...mid)的范围内查找
			r = mid
		}
	}
	return -1
}
