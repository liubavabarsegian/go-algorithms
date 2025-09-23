package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count, maxCount := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		}
		if nums[i] == 0 {
			count = 0
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) // 3
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1})) // 2
}
