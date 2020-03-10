package leetcode_golang

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	now := head.Next
	prev := head
	for {
		if now == nil {
			break
		}
		if now.Val == prev.Val {
			prev.Next = now.Next
			now = now.Next
		} else {
			prev = now
			now = now.Next
		}
	}
	return head
}
