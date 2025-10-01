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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead, evenHead := head, head.Next
	oddNode, evenNode := oddHead, evenHead

	for oddNode != nil && evenNode != nil && oddNode.Next != nil && evenNode.Next != nil {
		oddNode.Next = oddNode.Next.Next
		evenNode.Next = evenNode.Next.Next

		if oddNode.Next != nil {
			oddNode = oddNode.Next
		}
		if evenNode.Next != nil {
			evenNode = evenNode.Next
		}

	}
	oddNode.Next = evenHead
	return oddHead
}

// oddHead 1 evenHead 2
// 1 -> 2 -> 3 -> 4 -> 5
// 1 -> 3 -> 4 -> 5; 2 -> 4 -> 5
func main() {
	printLinkedList(oddEvenList(newLinkedList([]int{1, 2, 3, 4, 5})))       // [1,3,5,2,4]
	printLinkedList(oddEvenList(newLinkedList([]int{2, 1, 3, 5, 6, 4, 7}))) // [2,3,6,7,1,5,4]
	printLinkedList(oddEvenList(newLinkedList([]int{1})))                   // [1]
	printLinkedList(oddEvenList(newLinkedList([]int{})))                    // [1]
}
