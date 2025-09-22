package main

import (
	"fmt"
	"slices"
)

// already sorted
func twoSum(nums []int, target int) [][]int {
	start, end := 0, len(nums)-1
	res := make([][]int, 0)
	for start < end {
		if nums[start]+nums[end] > target {
			end--
			continue
		}
		if start > 0 && nums[start] == nums[start-1] {
			start++
			continue
		}
		if end < len(nums)-1 && nums[end] == nums[end+1] {
			end--
			continue
		}

		if nums[start]+nums[end] == target {
			res = append(res, []int{nums[start], nums[end]})
		}
		start++
	}
	return res
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoSumRes := twoSum(nums[i+1:], nums[i]*-1)
		for _, v := range twoSumRes {
			v = append([]int{nums[i]}, v...)
			res = append(res, v)
		}

	}
	return res
}

func main() {
	// -4 -1 -1 0 1 2
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [0, 0, 0]
	fmt.Println(threeSum([]int{0, 0, 0, 0}))          // [0, 0, 0, 0]
}
