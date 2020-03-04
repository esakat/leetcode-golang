package leetcode_golang

import (
	"log"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := 6
	actual := maxSubArray(input)
	if expect != actual {
		log.Fatalf("expect: %d, got: %d", expect, actual)
	}
}
