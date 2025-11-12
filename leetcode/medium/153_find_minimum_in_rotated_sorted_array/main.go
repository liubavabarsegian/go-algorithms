package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	i := len(nums) / 2
	for {
		if i < len(nums)-1 && nums[i] > nums[i+1] {
			return nums[i+1]
		} else if i < len(nums)-1 && nums[i] > nums[0] {
			i += 1
		} else {
			i -= 1
		}
	}
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // 11
	fmt.Println(findMin([]int{1}))                   // 1
	fmt.Println(findMin([]int{2, 1}))                // 1
	fmt.Println(findMin([]int{2, 3, 4, 5, 1}))       // 1

}
