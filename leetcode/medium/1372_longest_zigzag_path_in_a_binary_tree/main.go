package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFSZigZag(root *TreeNode, next string, counter int) int {
	if root == nil {
		return counter
	}

	counter++

	switch next {
	case "right":
		return max(DFSZigZag(root.Right, "left", counter), DFSZigZag(root.Left, "right", 0))
	case "left":
		return max(DFSZigZag(root.Left, "right", counter), DFSZigZag(root.Right, "left", 0))
	}
	return 0
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var left, right int
	if root.Left != nil {
		left = DFSZigZag(root.Left, "right", 0)
	}
	if root.Right != nil {
		right = DFSZigZag(root.Right, "left", 0)
	}

	return max(left, right)
}
