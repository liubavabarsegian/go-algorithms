package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftTree := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(leftTree)
	return root
}
