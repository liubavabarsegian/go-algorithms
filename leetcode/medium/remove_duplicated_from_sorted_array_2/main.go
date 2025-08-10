package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 0

	for _, v := range nums {
		if k < 2 || nums[k-2] != v {
			nums[k] = v
			k++
		}
	}
	return k
}

func removeDuplicates2(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}

	k := 2
	for i := 2; i < len(nums); i++ {
		if k < 2 || nums[k-2] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
