package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}

	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		min := minNode(root.Right)
		root.Val = min.Val
		root.Right = deleteNode(root.Right, min.Val)

	}
	return root
}

func main() {
	fmt.Println(deleteNode(
		&TreeNode{Val: 3,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 3},
		}, 2))
}
