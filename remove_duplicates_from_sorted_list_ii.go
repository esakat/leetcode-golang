package leetcode_golang

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[int]int)
	for {
		if head == nil {
			break
		}
		m[head.Val]++
		head = head.Next
	}

	var l, ans *ListNode
	sortInt := []int{}
	for k, v := range m {
		if v == 1 {
			sortInt = append(sortInt, k)
		}
	}
	sort.Ints(sortInt)
	for _, v := range sortInt {
		n := ListNode{
			Val:  v,
			Next: nil,
		}
		if l == nil {
			l = &n
			ans = &n
		} else {
			l.Next = &n
			l = l.Next
		}
	}
	return ans
}
