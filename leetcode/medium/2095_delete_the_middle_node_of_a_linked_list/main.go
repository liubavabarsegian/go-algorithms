package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newLinkedList(nodes []int) *ListNode {
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

// Алгоритм быстрого и медленного указателя
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	prev := &ListNode{}
	prev.Next = head
	slow, fast := prev, prev.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	printLinkedList(deleteMiddle(newLinkedList([]int{1, 3, 4, 7, 1, 2, 6}))) // delete 7
	printLinkedList(deleteMiddle(newLinkedList([]int{1, 2, 3, 4})))          // delete 3
	printLinkedList(deleteMiddle(newLinkedList([]int{2, 1})))                // delete 1
	printLinkedList(deleteMiddle(newLinkedList([]int{2})))                   // delete 2
}
