package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "(([[]]))"
	fmt.Println(s, " valid is ", isValid(s))
}
