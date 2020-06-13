package array

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {

	nums := []int{1, 2, 3, 1}
	k := 3

	fmt.Println(containsNearbyDuplicate(nums, k))

}
