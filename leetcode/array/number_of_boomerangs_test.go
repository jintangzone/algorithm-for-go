package array

import (
	"fmt"
	"testing"
)

func TestNumberOfBoomerangs(t *testing.T) {

	points := [][]int{
		{0, 0},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	fmt.Println(numberOfBoomerangs(points))

}
