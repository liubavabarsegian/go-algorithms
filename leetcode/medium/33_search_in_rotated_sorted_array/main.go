package main

import "fmt"

func search(nums []int, target int) int {
	i := 0
	for i < len(nums)-1 && nums[i] < nums[i+1] {
		i++
	}

	shift := len(nums[i+1:])
	var start, end int
	if target >= nums[0] {
		start = 0
		end = len(nums) - shift - 1
	} else {
		start = len(nums) - shift
		end = len(nums) - 1
	}

	for start <= end {
		medium := (start + end) / 2

		if nums[medium] == target {
			return medium
		} else if nums[medium] > target {
			end = medium - 1
		} else {
			start = medium + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(search([]int{1}, 0))                   // -1
	fmt.Println(search([]int{4, 5, 6, 1}, 1))          // 3
}
