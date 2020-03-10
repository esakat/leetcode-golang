package leetcode_golang

func maxDepth(root *TreeNode) int {
	return getDepth(root, 0)
}

func getDepth(n *TreeNode, depth int) int {
	if n == nil {
		return depth
	}
	a := getDepth(n.Left, depth+1)
	b := getDepth(n.Right, depth+1)
	if a > b {
		return a
	} else {
		return b
	}
}
