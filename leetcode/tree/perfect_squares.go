package tree

type pNode struct {
	Val   int
	Level int
}

func NewPNode(v, l int) pNode {
	return pNode{Val: v, Level: l}
}

func numSquares(n int) int {
	queue := make([]pNode, 0)
	visited := make(map[int]bool)
	queue = append(queue, NewPNode(n, 0))
	visited[n] = true
	for len(queue) > 0 {
		size := len(queue)
		frontNode := queue[0]
		queue = queue[1:size]
		if frontNode.Val == 0 {
			return frontNode.Level
		}
		i := 1
		for {
			le := frontNode.Val - i*i
			if le < 0 {
				break
			}
			if le == 0 {
				return frontNode.Level + 1
			}
			if le > 0 {
				if _, ok := visited[le]; !ok {
					queue = append(queue, NewPNode(le, frontNode.Level+1))
				}
			}
			i++
		}
	}
	return 0
}
