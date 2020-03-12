package leetcode_golang

// https://leetcode.com/problems/merge-two-binary-trees/
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	t1 = merge(t1, t2)
	return t1
}

func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t2 == nil {
		return t1
	}
	if t1 == nil {
		t1 = &TreeNode{}
	}

	t1.Val += t2.Val

	if t1.Left == nil && t2.Left != nil {
		t1.Left = &TreeNode{}
	}
	if t1.Right == nil && t2.Right != nil {
		t1.Right = &TreeNode{}
	}
	merge(t1.Left, t2.Left)
	merge(t1.Right, t2.Right)
	return t1
}
