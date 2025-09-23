package main

import "fmt"

func longestSubarray(nums []int) int {
	count, deleted, maxCount := 0, false, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		}
		if nums[i] == 0 {
			if deleted && nums[i-count-1] == 1 {
				for nums[i-count-1] == 1 {
					count--
				}
			} else if !deleted {
				deleted = true
			}
		}
		maxCount = max(maxCount, count)
	}

	if len(nums) == maxCount {
		maxCount--
	}

	return maxCount
}

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                                                                                                          // 3
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))                                                                                           // 5
	fmt.Println(longestSubarray([]int{1, 1, 1}))                                                                                                             // 2
	fmt.Println(longestSubarray([]int{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1})) // 18
}
