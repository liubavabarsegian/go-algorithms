package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newLinkedList(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}
	head := &ListNode{Val: nodes[0]}
	copyHead := head
	for i := 1; i < len(nodes); i++ {
		newNode := &ListNode{Val: nodes[i]}
		copyHead.Next = newNode
		copyHead = copyHead.Next
	}
	return head
}

func printLinkedList(head *ListNode) {
	copyHead := head
	for copyHead != nil {
		fmt.Print(copyHead.Val, "->")
		copyHead = copyHead.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		realNext := current.Next
		current.Next = prev
		prev = current
		current = realNext
	}
	return prev
}

func main() {
	printLinkedList(reverseList(newLinkedList([]int{1, 2, 3, 4, 5})))
	printLinkedList(reverseList(newLinkedList([]int{1, 2})))
	printLinkedList(reverseList(newLinkedList([]int{1})))
	printLinkedList(reverseList(newLinkedList([]int{})))
}
