package leetcode_golang

import (
	"log"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	//expect := [][]int{}
	actual := combinationSum(candidates, target)
	log.Printf("answer: %v", actual)
}
