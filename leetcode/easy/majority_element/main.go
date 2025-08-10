package main

import "fmt"

func majorityElement(nums []int) int {
	num, counter := nums[0], 0

	for i := 0; i < len(nums); i++ {
		if counter == 0 {
			num = nums[i]
		}
		if num == nums[i] {
			counter++
		} else {
			counter--
		}
	}

	counter = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			counter++
		}
	}

	if counter > len(nums)/2 {
		return num
	}

	return -1
}

func main() {
	nums := []int{6, 6, 6, 7, 7}
	fmt.Println(majorityElement((nums)))
}
