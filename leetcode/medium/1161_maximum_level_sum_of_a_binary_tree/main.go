package main

import (
	"fmt"
	"math"
)

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
		q.elems = q.elems[1:]
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

func maxLevelSum(root *TreeNode) int {
	maxSum := math.MinInt
	level := 0
	maxLevel := 0

	queue := &Queue{elems: []*TreeNode{root}}

	for queue.Peek() != nil {
		levelSize := len(queue.elems)
		level++
		sum := 0

		for range levelSize {
			node := queue.Out()
			sum += node.Val

			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}

		}

		if sum > maxSum {
			maxSum = sum
			maxLevel = level
		}
	}

	return maxLevel
}

func main() {
	fmt.Println(maxLevelSum(
		&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		}))

	fmt.Println(maxLevelSum(
		&TreeNode{Val: 1,
			Left: &TreeNode{Val: 7,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: -8},
			},
			Right: &TreeNode{Val: 0},
		}))

	fmt.Println(maxLevelSum(
		&TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}))
}
