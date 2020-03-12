package leetcode_golang

// TODO: マイナスの値が含まれているのこの方法だとだめ
// https://leetcode.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	l, r := 0, 0
	up_r := true
	ans := 0
	now := nums[0]
	for {
		if len(nums) == l {
			break
		}
		if now == k {
			ans++
		} else if now > k {
			up_r = false
		} else {
			up_r = true
		}

		if len(nums)-1 <= r {
			up_r = false
		}

		if up_r {
			r++
			now += nums[r]
		} else {
			now -= nums[l]
			l++
		}
	}

	return ans
}
