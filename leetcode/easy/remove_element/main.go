package main

import "fmt"

func removeElement(nums []int, val int) int {
	k, i := 0, 0
	size := len(nums)

	for i < size {
		for i < size && nums[i] == val {
			i++
		}

		if i < size {
			nums[k] = nums[i]
			k++
		}
		i++
	}

	return k
}

func removeElement2(nums []int, val int) int {
	k := 0

	for i, _ := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)

	nums = []int{3, 2, 2, 3}
	fmt.Println(removeElement2(nums, 3))
	fmt.Println(nums)
}
