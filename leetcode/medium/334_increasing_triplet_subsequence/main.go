package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	res := make([]int, 0, 3)

	res = append(res, nums[0])
	potentialNewMin := nums[0]
	for i := 1; i < len(nums); i++ {
		// записываем в массив если элемент болше последнего
		if len(res) == 3 {
			return true
		}

		if nums[i] > res[len(res)-1] {
			res = append(res, nums[i])
		} else if nums[i] > potentialNewMin {
			res = res[0:0]
			res = append(res, potentialNewMin, nums[i])
		} else if nums[i] < nums[0] {
			potentialNewMin = nums[i]
		}
	}
	return len(res) >= 3
}

// 2, 1, 5, 0, 4, 6
// min = 2
// min = 1
// next = 5
// min = 0
// next = 6
// 20, 100, 10, 12, 5, 13
// min = 20
// next = 100
// min = 10
// next = 12
// potential min = 5
// next = 13
func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))           // true
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))           // false
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))        // true
	fmt.Println(increasingTriplet([]int{}))                        // false
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))  // true
	fmt.Println(increasingTriplet([]int{1, 2, 3, 0, 4}))           // true
	fmt.Println(increasingTriplet([]int{9, 10, 5, 11, 10, 9, 8}))  // true
	fmt.Println(increasingTriplet([]int{-4, -5, -1, 5, 3, 2, -3})) // true
}
