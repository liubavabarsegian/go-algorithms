package main

import (
	"fmt"
)

func findMaxAverage1(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxAverage := float64(sum) / float64(k)

	for i := 0; i <= len(nums)-k; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += nums[j]
		}
		maxAverage = max(maxAverage, float64(sum)/float64(k))
	}
	return maxAverage
}

// sliding window
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxAverage := float64(sum) / float64(k)

	for i := 0; i < len(nums)-k; i++ {
		sum += nums[i+k] - nums[i]
		maxAverage = max(maxAverage, float64(sum)/float64(k))
	}
	return maxAverage
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
	fmt.Println(findMaxAverage([]int{5}, 1))                    //  5
	fmt.Println(findMaxAverage([]int{5}, 2))                    //  0
	fmt.Println(findMaxAverage([]int{0, 1, 1, 3, 3}, 4))        //  2.0
}
