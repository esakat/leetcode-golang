package leetcode_golang

import (
	"log"
	"reflect"
	"sort"
)

// DP Solution
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target == 0 {
		return [][]int{[]int{}}
	}

	dp := make([][][]int, target+1)
	for i := 1; i <= target; i++ {
		log.Println("----")
		for _, v := range candidates {
			if i-v < 0 {
				continue
			} else if i == v {
				log.Printf("%d", i)
				dp[i] = append(dp[i], []int{v})
			} else if len(dp[i-v]) > 0 {
				log.Printf("%d, %v", i-v, dp[i-v])
				for _, t := range dp[i-v] {
					s := make([]int, len(t))
					copy(s, t)
					s = append(s, v)
					dp[i] = append(dp[i], s)
				}
			}
		}
		log.Printf("%d, %v", i, dp[i])
	}

	ans := [][]int{}

	for _, v := range dp[target] {
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })

		flag := true
		for _, t := range ans {
			if reflect.DeepEqual(t, v) {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, v)
		}
	}

	return ans
}
