package leetcode_golang

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -2147483648
	}
	tmp := -999999999999
	ans := tmp
	for _, v := range nums {
		if tmp+v > v {
			tmp += v
		} else {
			tmp = v
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}
