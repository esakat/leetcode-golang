package leetcode_golang

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	ans := 0
	determined := 0
	i := 0
	for {
		if i > len(nums)-1 {
			break
		}
		if i == len(nums)-1 {
			ans += nums[i]
			tmp := i - 2
			for {
				if tmp < determined {
					break
				}
				ans += nums[tmp]
				tmp -= 2
			}
			break
		}
		if nums[i] > nums[i+1] {
			ans += nums[i]
			tmp := i - 2
			for {
				if tmp < determined {
					break
				}
				ans += nums[tmp]
				tmp -= 2
			}
			i += 2
			determined = i
		} else {
			i++
		}
	}

	return ans
}
