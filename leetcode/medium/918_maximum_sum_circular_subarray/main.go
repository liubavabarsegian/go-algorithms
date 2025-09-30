package main

import "fmt"

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func maxSubarraySumCircular(nums []int) int {
// 	circle := append(nums, nums...)
// 	totalSum := circle[0]

// 	currentSum := sum(circle)
// 	for len(circle) > 0 {
// 		if circle[0] > circle[len(circle)-1] {
// 			currentSum -= circle[len(circle)-1]
// 			circle = circle[:len(circle)-1]
// 		} else {
// 			currentSum -= circle[0]
// 			circle = circle[1:]

// 		}
// 		if len(circle) <= len(nums) {
// 			totalSum = max(totalSum, currentSum)
// 		}
// 		fmt.Println(circle, totalSum, currentSum)
// 	}
// 	return totalSum
// }

func maxSubarraySumCircular(nums []int) int {
	circle := append(nums, nums...)

	maxSubarray := circle[0]
	newMax := circle[0]
	subArrayLen := 0

	for i := 1; i < len(circle); i++ {
		fmt.Println(maxSubarray, newMax, circle)
		offset := i % len(circle)
		// если число положительное, то оно увеличит сумму, поэтому его берем
		for subArrayLen >= len(nums)-1 || offset > subArrayLen-1 && newMax-circle[offset-subArrayLen] < newMax {

			newMax -= circle[offset-subArrayLen]
			fmt.Println("new max", newMax, circle[offset-subArrayLen+1], offset)
			subArrayLen--
		}
		if newMax < 0 || offset > 0 && circle[offset-1] < 0 && newMax+circle[offset] < newMax {
			// сброс подмассива
			fmt.Println("reset", newMax, circle[offset])
			newMax = circle[offset]
			subArrayLen = 0
		} else {
			newMax += circle[offset]
			subArrayLen++
			fmt.Println("add", newMax, circle[offset])
		}

		maxSubarray = max(maxSubarray, newMax)
	}
	return maxSubarray
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2})) // 3
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))     // 10
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))   // -2
	fmt.Println(maxSubarraySumCircular([]int{3, -1, 2, -1})) // 4
}
