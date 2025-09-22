package main

import (
	"fmt"
	"slices"
)

// n2 solution (if i can't sort)
func countPairs1(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}

func countPairs(nums []int, target int) int {
	slices.Sort(nums)
	count, start, end := 0, 0, len(nums)-1

	for start < end {
		// сдвигаем правый указатель до тех пор, пока минимальное значение + текущее не удовлетворяют условию
		if nums[start]+nums[end] >= target {
			end--
		} else {
			// если мы попали сюда, значит сумма минимального и максимального числа меньше target
			// значит остальные элементы до end в сумме с минимальным будут меньше target
			// поэтому мы сразу считаем их в count
			count += (end - start)
			start++
		}
	}
	return count
}

func main() {
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))           // 3
	fmt.Println(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2)) // 10
	fmt.Println(countPairs([]int{}, -2))                        // 0
}
