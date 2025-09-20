package main

import "fmt"

func moveZeroes1(nums []int) {
	for i := range nums {
		nonZeroIter := i
		if nums[i] == 0 {
			for nonZeroIter < len(nums)-1 && nums[nonZeroIter] == 0 {
				nonZeroIter++
			}
			nums[i], nums[nonZeroIter] = nums[nonZeroIter], nums[i]
		}
	}
}

// time optimization
func moveZeroes(nums []int) {
	nonZeroIter := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[nonZeroIter] = nums[i]
			nonZeroIter++
		}
	}

	for i := nonZeroIter; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) //1,3,12,0,0,0

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums) //0

	nums = []int{0, 0, 0, 0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums) //1, 0, 0, 0, 0

	nums = []int{1, 0}
	moveZeroes(nums)
	fmt.Println(nums) //1, 0, 0, 0, 0
}
