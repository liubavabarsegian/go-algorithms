package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func getLeaves(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leaves := make([]int, 0)
	if root.Left == nil && root.Right == nil {
		leaves = append(leaves, root.Val)
	}

	leaves = append(leaves, getLeaves(root.Left)...)
	leaves = append(leaves, getLeaves(root.Right)...)

	return leaves
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := getLeaves(root1)
	leaves2 := getLeaves(root2)
	fmt.Println(leaves1, leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(leafSimilar(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}))
	fmt.Println(leafSimilar(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}))
}
