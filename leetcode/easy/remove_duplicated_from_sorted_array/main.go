package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 0

	for i, v := range nums {
		if i == 0 || nums[i-1] != v {
			nums[k] = v
			k++
		}
	}
	return k
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
