package leetcode_golang

func dfs_tree(node *TreeNode, m map[int][]int, depth int) {
	if _, ok := m[depth]; ok {
		m[depth] = append(m[depth], node.Val)
	} else {
		m[depth] = []int{node.Val}
	}

	if node.Left != nil {
		dfs_tree(node.Left, m, depth+1)
	}

	if node.Right != nil {
		dfs_tree(node.Right, m, depth+1)
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)

	if root != nil {
		dfs_tree(root, m, 1)
	}

	ans := [][]int{}

	depth := 1
	for {
		if val, ok := m[depth]; ok {
			if depth%2 == 0 {
				// reverse
				for i, j := 0, len(val)-1; i < j; i, j = i+1, j-1 {
					val[i], val[j] = val[j], val[i]
				}
			}
			ans = append(ans, val)
		} else {
			break
		}
		depth++
	}

	return ans
}
