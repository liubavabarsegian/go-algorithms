package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var currentNode *ListNode = &ListNode{}
	var dummyHead *ListNode = currentNode

	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		if sum > 9 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}

		currentNode.Val = sum

		if l1 != nil || l2 != nil || carry != 0 {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next
		}
	}

	return dummyHead
}

func printList(list *ListNode) {
	pointer := list

	for pointer != nil {
		fmt.Print(pointer.Val, "-->")
		pointer = pointer.Next
	}
	fmt.Println()
}

func main() {
	l1 := ListNode{9, &ListNode{9, nil}}
	l2 := ListNode{9, &ListNode{9, &ListNode{9, nil}}}

	result := addTwoNumbers(&l1, &l2)
	printList(result)
}
