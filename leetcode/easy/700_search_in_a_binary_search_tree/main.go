package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {

	for root != nil {
		switch {
		case val == root.Val:
			return root
		case val > root.Val:
			root = root.Right
		case val < root.Val:
			root = root.Left
		}
	}

	return nil
}

func main() {
	fmt.Println(searchBST(
		&TreeNode{Val: 3,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 3},
		}, 2))
}
