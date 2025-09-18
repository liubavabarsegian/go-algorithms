package main

import "fmt"

func productExceptSelf(nums []int) []int {
	results := make([]int, len(nums))

	product := 1
	for i, v := range nums {
		results[i] = product

		product = product * v
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		results[i] = results[i] * product

		product = product * nums[i]
	}
	return results
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(arr))

	arr = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(arr))
}
