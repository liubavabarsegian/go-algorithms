package main

import (
	"fmt"
	"slices"
)

func maxOperations(nums []int, k int) int {
	slices.Sort(nums)
	start, end, count := 0, len(nums)-1, 0

	for start < end {
		if nums[start]+nums[end] > k {
			end--
			continue
		}
		if nums[start]+nums[end] < k {
			start++
			continue
		}
		if nums[start]+nums[end] == k {
			count++
			start++
			end--
		}
	}
	return count
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))    // 2
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6)) // 1
	fmt.Println(maxOperations([]int{}, 1))              // 0
	fmt.Println(maxOperations([]int{3, 5, 1, 5}, 2))    // 0
}
