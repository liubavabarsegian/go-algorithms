package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 1
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return count + max(left, right)
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(maxDepth(root))

	root = &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println(maxDepth(root))

	fmt.Println(maxDepth(nil))
}
