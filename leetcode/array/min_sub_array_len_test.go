package array

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}
