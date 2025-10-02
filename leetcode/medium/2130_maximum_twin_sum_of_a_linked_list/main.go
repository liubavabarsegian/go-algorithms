package main

import (
	"fmt"
	"math"
)

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
	copyHead := head
	var prev *ListNode

	for copyHead != nil {
		realNex := copyHead.Next
		copyHead.Next = prev
		prev = copyHead
		copyHead = realNex
	}
	return prev
}

func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}

	slow, fast, copyHead := head, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = reverseList(slow)
	maxPairSum := math.MinInt
	for slow != nil {
		maxPairSum = max(copyHead.Val+slow.Val, maxPairSum)
		slow = slow.Next
		copyHead = copyHead.Next
	}
	return maxPairSum
}

// -1 5
// 5 2
// 4 nil
func main() {
	fmt.Println(pairSum(newLinkedList([]int{5, 4, 2, 1})))      //6
	fmt.Println(pairSum(newLinkedList([]int{4, 2, 3, 3})))      //7
	fmt.Println(pairSum(newLinkedList([]int{1, 10000})))        //10001
	fmt.Println(pairSum(newLinkedList([]int{})))                //0
	fmt.Println(pairSum(newLinkedList([]int{4, 2, 10, 1})))     //12
	fmt.Println(pairSum(newLinkedList([]int{4, 2, 3})))         //7
	fmt.Println(pairSum(newLinkedList([]int{-1, -3, -10, -3}))) //-4
	fmt.Println(pairSum(newLinkedList([]int{-1})))              //-1
}
