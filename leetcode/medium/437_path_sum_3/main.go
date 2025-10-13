package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFSsubarraySum(root *TreeNode, prefixSum int, frequenciesMap map[int]int, targetSum int) int {
	if root == nil {
		return 0
	}
	result := 0

	prefixSum += root.Val
	// we update the result only when prefixSum - k exists in the map
	if _, ok := frequenciesMap[prefixSum-targetSum]; ok {
		result += frequenciesMap[prefixSum-targetSum]
	}
	if prefixSum == targetSum {
		result++
	}
	frequenciesMap[prefixSum]++

	left := DFSsubarraySum(root.Left, prefixSum, frequenciesMap, targetSum)
	right := DFSsubarraySum(root.Right, prefixSum, frequenciesMap, targetSum)

	// backtracking. это эффективнее копирования
	// просто удаляем нижний элемент из мапы после обработки или уменьшаем счетчик
	frequenciesMap[prefixSum]--

	return result + left + right
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	frequenciesMap := make(map[int]int)
	prefixSum := 0

	return DFSsubarraySum(root, prefixSum, frequenciesMap, targetSum)
}

func main() {
	fmt.Println(pathSum(
		&TreeNode{Val: 10,
			Left: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 3,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: -2},
				},
				Right: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{Val: -3,
				Right: &TreeNode{Val: 11,
					Right: &TreeNode{Val: -3}},
			},
		}, 8))

	//[5,4,8,11,null,13,4,7,2,null,null,5,1]
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
					Right: &TreeNode{Val: 1}},
			},
		}, 22))
}
