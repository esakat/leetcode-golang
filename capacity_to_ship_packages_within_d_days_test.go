package leetcode_golang

import (
	"log"
	"testing"
)

func Test_shipWithinDays(t *testing.T) {
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := 5
	expect := 15
	actual := shipWithinDays(candidates, d)
	if expect != actual {
		log.Fatalf("expect: %v, actual: %v", expect, actual)
	}
}
