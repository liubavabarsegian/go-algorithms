package main

import (
	"fmt"
	"math/big"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getNumber(list *ListNode) *big.Int {
	num := big.NewInt(0)
	multiplier := big.NewInt(1)
	step := big.NewInt(10)

	pointer := list

	for pointer != nil {
		nodeValue := big.NewInt(int64(pointer.Val))
		sumValue := new(big.Int).Mul(nodeValue, multiplier)

		num.Add(num, sumValue)

		pointer = pointer.Next
		multiplier.Mul(multiplier, step)
	}

	return num
}

func numberToList(number *big.Int) *ListNode {
	var prevNode *ListNode = nil
	var headNode *ListNode = nil

	if number.Cmp(big.NewInt(0)) == 0 {
		return &ListNode{0, nil}
	}

	for number.Cmp(big.NewInt(0)) > 0 {
		digit := new(big.Int).Mod(number, big.NewInt(10)).Int64()
		currentNode := ListNode{int(digit), nil}

		if prevNode == nil {
			prevNode = &currentNode
			headNode = &currentNode
		} else {
			prevNode.Next = &currentNode
			prevNode = &currentNode
		}
		number.Div(number, big.NewInt(10))
	}

	return headNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := getNumber(l1)
	num2 := getNumber(l2)

	result := new(big.Int).Add(num1, num2)

	return numberToList(result)
}

func main() {
	l1 := ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	l3 := addTwoNumbers(&l1, &l2)
	for l3 != nil {
		fmt.Print(l3.Val, "-->")
		l3 = l3.Next
	}
}
