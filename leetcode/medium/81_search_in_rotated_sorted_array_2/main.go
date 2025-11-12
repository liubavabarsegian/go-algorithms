package main

import "fmt"

func search(nums []int, target int) bool {
	i := 0
	for i < len(nums)-1 && nums[i] <= nums[i+1] {
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
			return true
		} else if nums[medium] > target {
			end = medium - 1
		} else {
			start = medium + 1
		}
	}

	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0)) // true
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3)) // false
	fmt.Println(search([]int{2, 2, 2, 3, 2, 2, 2}, 3)) // true
}
