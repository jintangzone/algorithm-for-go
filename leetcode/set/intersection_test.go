package set

import (
	"fmt"
	"testing"
)

//
// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2]
//
//
// 示例 2:
//
// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [9,4]
//
// 说明:
//
//
// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。
//
// Related Topics 排序 哈希表 双指针 二分查找

func TestIntersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
