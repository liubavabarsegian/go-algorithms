package main

import "fmt"

// первое решение по интуиции
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	start, end := 0, 0
	maxWindow, maxWindowIndex := nums[0], 0

	maxWindows := make([]int, 0)
	for end < k {
		if nums[end] >= maxWindow {
			maxWindow = nums[end]
			maxWindowIndex = end
		}
		end++
	}
	start++
	maxWindows = append(maxWindows, maxWindow)

	for end < len(nums) {
		// если максимальный элемент вышел за пределы окна
		if maxWindowIndex < start {
			maxWindow, maxWindowIndex = nums[start], start
			for i := start + 1; i <= end; i++ {
				if nums[i] >= maxWindow {
					maxWindow = nums[i]
					maxWindowIndex = i
				}
			}
		} else if nums[end] >= maxWindow {
			maxWindow = nums[end]
			maxWindowIndex = end
		}

		maxWindows = append(maxWindows, maxWindow)
		start++
		end++
	}
	return maxWindows
}

// func maxSlidingWindow(nums []int, k int) []int {

// }

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // [3,3,5,5,6,7]
	fmt.Println(maxSlidingWindow([]int{1}, 1))                        // [1]
}
