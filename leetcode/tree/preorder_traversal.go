package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Cmd struct {
	s    string
	node *TreeNode
}

func NewCmd(s string, node *TreeNode) Cmd {
	return Cmd{s: s, node: node}
}

func preorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := make([]Cmd, 0)
	stack = append(stack, NewCmd("go", root))

	for len(stack) > 0 {
		n := len(stack)
		cmd := stack[n-1]
		stack = stack[:n-1]
		if cmd.s == "print" {
			res = append(res, cmd.node.Val)
		} else {
			if cmd.node.Right != nil {
				stack = append(stack, NewCmd("go", cmd.node.Right))
			}
			if cmd.node.Left != nil {
				stack = append(stack, NewCmd("go", cmd.node.Left))
			}
			stack = append(stack, NewCmd("print", cmd.node))
		}
	}
	return res
}
