package leetcode_golang

func isValidBST(root *TreeNode) bool {
	return reccursive(root, -1<<63+1, 1<<63-1)
}

func reccursive(n *TreeNode, min, max int64) bool {
	if n == nil {
		return true
	}
	if int64(n.Val) <= min || int64(n.Val) >= max {
		return false
	}
	return reccursive(n.Left, min, int64(n.Val)) && reccursive(n.Right, int64(n.Val), max)
}
