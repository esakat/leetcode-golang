package leetcode_golang

import (
	"sort"
)

func dfs(index, target, sum int, ans, candidates []int, res *[][]int) {
	if sum == target {
		*res = append(*res, ans)
	} else if sum < target {
		// すでにindexより前の要素を使ったのは確かめてるので、不要
		// つまり インデックスでの組み合わせで 0,0,1　が確定した後に 0,1,0 を探索しないように 0,1 ときたら見るのは 0,1,1 0,1,2のようになる
		for i := index; i < len(candidates); i++ {
			tmp := sum + candidates[i]
			if tmp > target {
				break
			}
			s := make([]int, len(ans))
			copy(s, ans)
			s = append(s, candidates[i])
			dfs(i, target, tmp, s, candidates, res)
		}
	}
}

// DFS pattern
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target == 0 {
		return [][]int{[]int{}}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	res := [][]int{}

	for i, v := range candidates {
		dfs(i, target, v, []int{v}, candidates, &res)
	}

	return res
}

/** DP Solution
計算時間すごくかかる
*/
//func combinationSum(candidates []int, target int) [][]int {
//	if len(candidates) == 0 || target == 0 {
//		return [][]int{[]int{}}
//	}
//
//	dp := make([][][]int, target+1)
//	for i := 1; i <= target; i++ {
//		log.Println("----")
//		for _, v := range candidates {
//			if i-v < 0 {
//				continue
//			} else if i == v {
//				log.Printf("%d", i)
//				dp[i] = append(dp[i], []int{v})
//			} else if len(dp[i-v]) > 0 {
//				log.Printf("%d, %v", i-v, dp[i-v])
//				for _, t := range dp[i-v] {
//					s := make([]int, len(t))
//					copy(s, t)
//					s = append(s, v)
//					dp[i] = append(dp[i], s)
//				}
//			}
//		}
//		log.Printf("%d, %v", i, dp[i])
//	}
//
//	ans := [][]int{}
//
//	for _, v := range dp[target] {
//		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
//
//		flag := true
//		for _, t := range ans {
//			if reflect.DeepEqual(t, v) {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			ans = append(ans, v)
//		}
//	}
//
//	return ans
//}
