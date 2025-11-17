package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	nonEmptyListIter := 0
	for nonEmptyListIter < len(lists)-1 {
		if lists[nonEmptyListIter] != nil {
			break
		}
		nonEmptyListIter++
	}

	head := lists[nonEmptyListIter]
	for i := nonEmptyListIter + 1; i < len(lists); i++ {
		current := lists[i]
		mergedCurrent := head

		for current != nil {
			node := &ListNode{Val: current.Val}

			for mergedCurrent != nil && mergedCurrent.Next != nil && mergedCurrent.Next.Val <= node.Val {
				mergedCurrent = mergedCurrent.Next
			}

			// fmt.Println("node", node)
			// printList(mergedCurrent)
			if head == nil {
				head = node
			} else if mergedCurrent != nil {
				if mergedCurrent.Val > node.Val {
					node.Next = mergedCurrent
					head = node
					mergedCurrent = head
				} else {
					next := mergedCurrent.Next
					mergedCurrent.Next = node
					node.Next = next
				}
			}
			current = current.Next
		}

	}
	return head
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
	printList(mergeKLists([]*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 2, Next: &ListNode{Val: 6}},
	}))

	printList(mergeKLists([]*ListNode{
		&ListNode{Val: 1},
		&ListNode{Val: 0}},
	))

	printList(mergeKLists([]*ListNode{
		nil,
		&ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 11}}},
		nil,
		&ListNode{Val: 6, Next: &ListNode{Val: 10}},
	}))

	printList(mergeKLists([]*ListNode{
		nil,
	}))

	printList(mergeKLists([]*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}},
		&ListNode{Val: -4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}},
	}))
}
