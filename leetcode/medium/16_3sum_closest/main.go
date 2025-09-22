package main

import (
	"fmt"
	"math"
	"slices"
)

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}

func twoSum(nums []int, target int) int {
	slices.Sort(nums)
	start, end := 0, len(nums)-1
	diff := math.MaxInt

	for start < end {
		if abs(target-nums[start]-nums[end]) < abs(diff) {
			diff = target - nums[start] - nums[end]
		}

		// пользуемся тем, что у нас отсортированный массив
		// если текущая сумма меньше target, то сдвигаем левый указатель к большим числам
		// если текущая сумма больше target, то уменьшаем сумму, сдвигая правый указатель к меньшим числам
		// иначе (полное совпадние) - возвращаем target
		if nums[start]+nums[end] < target {
			start++
		} else if nums[start]+nums[end] > target {
			end--
		} else {
			return target
		}
	}
	return target - diff
}

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	diff := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		twoSum := twoSum(nums[i+1:], target-nums[i])
		if abs(target-twoSum-nums[i]) < abs(diff) {
			diff = target - twoSum - nums[i]
		}
	}
	return target - diff
}

// -4 -1 1 2, t = 1

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))      // 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))           // 0
	fmt.Println(threeSumClosest([]int{4, 6, 4}, 1))           // 14
	fmt.Println(threeSumClosest([]int{-4, 2, 2, 3, 3, 3}, 0)) // 0
	fmt.Println(twoSum([]int{-4, 2, 2, 3, 3, 3}, 4))          // 4
}
