package main

import "fmt"

// [2, 3, 1, 1, 4]
// 2 -> 3 or 2 -> 1
// base case:
// 1. return true if it is the last element
// 2. return false if did not achieve the last element
// not optimized solution
func canJump1(nums []int) bool {
	var result bool
	for i, maxJump := range nums {
		if i == len(nums)-1 {
			return true
		}

		for j := 1; j <= maxJump; j++ {
			result = canJump1(nums[i+j:])
			if result == true {
				return true
			}
		}
		if maxJump == 0 {
			return false
		}
	}
	return false
}

// [2, 3, 1, 1, 4]
func canJump2(nums []int) bool {
	gas := 0
	for _, maxJump := range nums {
		if gas < 0 {
			return false
		} else if maxJump > gas {
			gas = maxJump
		}
		gas--
	}
	return true
}

func canJump(nums []int) bool {
	reach := 0
	for i, maxJump := range nums {
		if reach < i {
			return false
		}

		if i+maxJump > reach {
			reach = i + maxJump
		}

		if i == len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))

	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))

	nums = []int{2, 5, 0, 0}
	fmt.Println(canJump(nums))

	nums = []int{2, 1, 0}
	fmt.Println(canJump(nums))

	nums = []int{2, 0, 0}
	fmt.Println(canJump(nums))

	nums = []int{2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6}
	fmt.Println(canJump(nums))
}
