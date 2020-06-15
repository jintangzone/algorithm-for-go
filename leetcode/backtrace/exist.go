package backtrace

type Board struct {
	direct  [4][2]int
	m, n    int
	visited [][]bool
}

func NewBoard(m, n int) *Board {
	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}
	return &Board{
		direct:  [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
		m:       m,
		n:       n,
		visited: visited,
	}
}

func (b *Board) inArea(x, y int) bool {
	return x >= 0 && x < b.m && y >= 0 && y < b.n
}

func (b *Board) searchWord(board [][]byte, word []byte, index, startx, starty int) bool {
	if index == len(word)-1 {
		return board[startx][starty] == word[index]
	}
	if board[startx][starty] == word[index] {
		b.visited[startx][starty] = true
		for i := 0; i < 4; i++ {
			newX := startx + b.direct[i][0]
			newY := starty + b.direct[i][1]
			if b.inArea(newX, newY) && !b.visited[newX][newY] && b.searchWord(board, word, index+1, newX, newY) {
				return true
			}
		}
		b.visited[startx][starty] = false
	}
	return false
}

func exist(board [][]byte, word string) bool {
	if len(board) <= 0 {
		return false
	}
	b := NewBoard(len(board), len(board[0]))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if b.searchWord(board, []byte(word), 0, i, j) {
				return true
			}
		}
	}
	return false
}
