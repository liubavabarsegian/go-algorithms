package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeHeight(root *TreeNode, height int) int {
	if root == nil {
		return height
	}

	height++
	return max(treeHeight(root.Left, height), treeHeight(root.Right, height))
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	balanced := abs(treeHeight(root.Left, 0), treeHeight(root.Right, 0)) <= 1
	if balanced && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}
