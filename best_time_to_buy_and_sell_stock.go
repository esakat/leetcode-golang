package leetcode_golang

func maxProfit(prices []int) int {
	maxVal := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 {
			maxVal[i] = prices[i]
		} else {
			if prices[i] > maxVal[i+1] {
				maxVal[i] = prices[i]
			} else {
				maxVal[i] = maxVal[i+1]
			}
		}
	}

	ans := 0
	for i := 0; i < len(prices); i++ {
		if maxVal[i]-prices[i] > ans {
			ans = maxVal[i] - prices[i]
		}
	}

	return ans
}
