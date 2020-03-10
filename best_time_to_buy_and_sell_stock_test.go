package leetcode_golang

import (
	"log"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	input := []int{}
	expect := 0
	actual := maxProfit(input)
	if expect != actual {
		log.Fatalf("expect: %d, got: %d", expect, actual)
	}
}
