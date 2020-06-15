package backtrace

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {

	//  ['A','B','C','E'],
	//  ['S','F','C','S'],
	//  ['A','D','E','E']

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "SEE"))

}
