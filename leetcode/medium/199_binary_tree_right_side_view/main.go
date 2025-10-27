package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	elems []*TreeNode
}

func (q *Queue) Out() *TreeNode {
	if len(q.elems) > 0 {
		elem := q.elems[0]
		q.elems = q.elems[1 : len(q.elems)-1]
		return elem
	}
	return nil
}

func (q *Queue) Peek() *TreeNode {
	if len(q.elems) > 0 {
		return q.elems[0]
	}
	return nil
}

func (q *Queue) Push(node *TreeNode) {
	q.elems = append(q.elems, node)
}

func BFS(res *[][]int, level int, root *TreeNode) {
	if root == nil {
		return
	}

	if len(*res) <= level {
		*res = append(*res, []int{})
		(*res)[level] = make([]int, 0)
	}

	(*res)[level] = append((*res)[level], root.Val)
	BFS(res, level+1, root.Left)
	BFS(res, level+1, root.Right)
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	BFS(&res, 0, root)

	rightValues := make([]int, 0)
	for i := range res {
		rightValues = append(rightValues, res[i][len(res[i])-1])
	}
	return rightValues
}

func main() {
	fmt.Println(rightSideView(
		&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		}))
}
