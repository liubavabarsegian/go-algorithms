package main

import "fmt"

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func minSubarraySum(nums []int) int {
	minSum := nums[0]
	currentMinSum := minSum

	for i := 1; i < len(nums); i++ {
		currentMinSum += nums[i]
		if nums[i] < currentMinSum {
			currentMinSum = nums[i]
		}
		minSum = min(minSum, currentMinSum)
	}
	return minSum
}

func maxSubarraySumCircular(nums []int) int {
	return sum(nums) - minSubarraySum(nums)
}

// 1 -2 3 2
func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2})) // 3
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))     // 10
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))   // -2
	fmt.Println(maxSubarraySumCircular([]int{3, -1, 2, -1})) // 4
}
