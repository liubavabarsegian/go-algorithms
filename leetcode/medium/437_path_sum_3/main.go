package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func paths(root *TreeNode, currentpath []int, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	currentpath = append(currentpath, root.Val)

	res := [][]int{}

	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return append(res, append([]int(nil), currentpath...))
	}

	res = append(res, paths(root.Left, append([]int(nil), currentpath...), targetSum-root.Val)...)
	res = append(res, paths(root.Right, append([]int(nil), currentpath...), targetSum-root.Val)...)

	return res
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	if root == nil {
		return [][]int{}
	}

	return paths(root, []int{}, targetSum)
}

func main() {
	fmt.Println(pathSum(
		&TreeNode{Val: 5,
			Left: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 11,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 2},
				},
			},
			Right: &TreeNode{Val: 8,
				Left: &TreeNode{Val: 13},
				Right: &TreeNode{Val: 4,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 1},
				},
			},
		}, 22))
}
