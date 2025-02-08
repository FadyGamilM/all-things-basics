package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	curr := head
	fast, slow := curr, curr
	if fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	} else {
		return false
	}

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
		if fast.Val == slow.Val {
			return true
		}
	}

	return false
}
