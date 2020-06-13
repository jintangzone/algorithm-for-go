package array

import (
	"fmt"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeros(nums)
	fmt.Println(nums)
}
