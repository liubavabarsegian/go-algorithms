package main

import (
	"fmt"
)

// 2, 3, 1, 1, 4
func jump(nums []int) int {
	jumps, curr, far := 0, 0, 0

	for i := range len(nums) - 1 {
		fmt.Println("1: ", far, i, nums[i], i+nums[i], curr)
		far = max(far, i+nums[i])

		if curr == i {
			jumps++
			curr = far
		}
		if curr > len(nums)-1 {
			break
		}

	}

	return jumps
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))

	nums = []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums))

	nums = []int{1, 2}
	fmt.Println(jump(nums))
}
