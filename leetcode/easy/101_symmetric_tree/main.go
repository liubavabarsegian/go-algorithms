package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(queue *[][]interface{}, level int, root *TreeNode) {

	if len(*queue) <= level {
		*queue = append(*queue, []interface{}{})
		(*queue)[level] = make([]interface{}, 0)
	}

	if root == nil {
		(*queue)[level] = append((*queue)[level], nil)
		return
	} else {
		(*queue)[level] = append((*queue)[level], root.Val)
	}

	BFS(queue, level+1, root.Left)
	BFS(queue, level+1, root.Right)
}

func isSymmetric(root *TreeNode) bool {
	queue := [][]interface{}{}

	BFS(&queue, 0, root)
	for _, q := range queue {
		for i := 0; i < len(q)/2; i++ {
			if (q)[i] != (q)[len(q)-1-i] {
				return false
			}
		}
	}
	return true
}
