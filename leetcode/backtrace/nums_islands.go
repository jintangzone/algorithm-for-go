package backtrace

type Gird struct {
	direct  [4][2]int
	m, n    int
	visited [][]bool
}

func NewGrid(m, n int) *Gird {
	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}
	return &Gird{
		direct:  [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
		m:       m,
		n:       n,
		visited: visited,
	}
}

func (g *Gird) inArea(x, y int) bool {
	return x >= 0 && x < g.m && y >= 0 && y < g.n
}

func (g *Gird) dfs(grid [][]byte, x, y int) {
	g.visited[x][y] = true
	for i := 0; i < 4; i++ {
		newX := x + g.direct[i][0]
		newY := y + g.direct[i][1]
		if g.inArea(newX, newY) && !g.visited[newX][newY] && grid[newX][newY] == '1' {
			g.dfs(grid, newX, newY)
		}
	}
	return
}

func numIslands(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	g := NewGrid(m, n)

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !g.visited[i][j] {
				res++
				g.dfs(grid, i, j)
			}
		}
	}
	return res
}
