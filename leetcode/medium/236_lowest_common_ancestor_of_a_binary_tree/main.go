package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		fmt.Println("a", root)
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(lowestCommonAncestor(root, root, root.Left))

	root = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))

	root = &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}}}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right))

	root = &TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right))
}
