package funny

import (
	"fmt"
	"testing"
)

func TestDivideRedPackage(t *testing.T) {
	list := DivideRedPackage(100, 10)
	for i, v := range list {
		fmt.Printf("第%d个人抢到 %.2f \n", i+1, v)
	}
}
