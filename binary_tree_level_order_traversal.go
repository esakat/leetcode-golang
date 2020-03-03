package leetcode_golang

// https://leetcode.com/problems/binary-tree-level-order-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	val   *TreeNode
	level int
}
type Stack []Pair

func (s *Stack) Push(v Pair) {
	*s = append(*s, v)
}

func (s *Stack) Pop() Pair {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	st := Stack{}
	st.Push(Pair{root, 0})
	for {
		if st.Empty() {
			break
		}
		now := st.Pop()
		if len(result) < now.level+1 {
			result = append(result, []int{})
		}
		result[now.level] = append(result[now.level], now.val.Val)
		left := now.val.Left
		right := now.val.Right
		if right != nil {
			st.Push(Pair{right, now.level + 1})
		}
		if left != nil {
			st.Push(Pair{left, now.level + 1})
		}
	}
	return result
}
