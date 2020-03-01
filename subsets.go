package leetcode_golang

// https://leetcode.com/problems/subsets/
func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	for _, num := range nums {
		var tmp [][]int
		for _, r := range result {
			t := make([]int, len(r))
			copy(t, r)
			t = append(t, num)
			tmp = append(tmp, t)
		}
		result = append(result, tmp...)
	}
	return result
}
