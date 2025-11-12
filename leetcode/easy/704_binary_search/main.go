package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		medium := (left + right) / 2

		if nums[medium] == target {
			return medium
		} else if nums[medium] < target {
			left = medium + 1
		} else if nums[medium] > target {
			right = medium - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	} else {
		return -1
	}
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9)) // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2)) // -1
	fmt.Println(search([]int{}, 2))                   // -1
	fmt.Println(search([]int{2, 5}, 2))               // 0
}
