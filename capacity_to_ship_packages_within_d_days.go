package leetcode_golang

// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
func shipWithinDays(weights []int, D int) int {
	var l, r int
	for _, w := range weights {
		if w > l {
			l = w
		}
		r += w
	}
	for {
		if l > r {
			break
		}
		capacity := l + (r-l)/2
		if moreLessDays(weights, D, capacity) {
			r = capacity - 1
		} else {
			l = capacity + 1
		}
	}

	return l
}

func moreLessDays(weights []int, D, capacity int) bool {
	now := 0
	size := 1
	for _, w := range weights {
		if w > capacity {
			return false
		}
		if now+w > capacity {
			now = w
			size++
		} else {
			now += w
		}
		if size > D {
			return false
		}
	}
	return true
}
