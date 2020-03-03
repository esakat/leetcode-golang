package leetcode_golang

import (
	"log"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	candidates := []int{7, 3, 2}
	target := 18
	//expect := [][]int{}
	actual := combinationSum(candidates, target)
	log.Printf("answer: %v", actual)
}
