package leetcode_golang

import (
	"log"
	"testing"
)

func Test_subsets(t *testing.T) {
	input := []int{3, 2, 4, 1}
	//expect := [][]int{[]int{}}
	actual := subsets(input)
	log.Println(actual)
}
