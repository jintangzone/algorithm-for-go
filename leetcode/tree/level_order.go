package tree

type pair struct {
	node  *TreeNode
	level int
}

func NewPair(node *TreeNode, level int) pair {
	return pair{node: node, level: level}
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]pair, 0)
	queue = append(queue, NewPair(root, 0))
	for len(queue) > 0 {
		n := len(queue)
		pair := queue[0]
		queue = queue[1:n]
		if pair.level == len(res) {
			res = append(res, []int{})
		}
		res[pair.level] = append(res[pair.level], pair.node.Val)
		if pair.node.Left != nil {
			queue = append(queue, NewPair(pair.node.Left, pair.level+1))
		}
		if pair.node.Right != nil {
			queue = append(queue, NewPair(pair.node.Right, pair.level+1))
		}
	}
	return res
}
