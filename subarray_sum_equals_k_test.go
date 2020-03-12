package leetcode_golang

import (
	"testing"
)

func Test_subarraySum(t *testing.T) {
	candidates := []int{28, 54, 7, -70, 22, 65, -6}
	k := 100
	expect := 1
	actual := subarraySum(candidates, k)
	if expect != actual {
		//log.Fatalf("expect: %v, actual: %v", expect, actual)
	}
}
