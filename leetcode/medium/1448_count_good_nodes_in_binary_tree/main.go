package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFS(root *TreeNode, maxVal int) int {
	count := 0

	if root == nil {
		return 0
	}

	if root.Val >= maxVal {
		count++
		maxVal = root.Val
	}

	count += DFS(root.Left, maxVal)
	count += DFS(root.Right, maxVal)

	return count
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return DFS(root, root.Val)
}

func main() {
	fmt.Println(goodNodes(&TreeNode{Val: 3, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}}))
}
