package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	nodes := make(map[*ListNode]bool)

	current := head
	for current != nil {
		if nodes[current] {
			return true
		}
		nodes[current] = true
		current = current.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
