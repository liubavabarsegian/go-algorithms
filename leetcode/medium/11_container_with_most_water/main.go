package main

import "fmt"

func maxArea(height []int) int {
	start, end, step := 0, len(height)-1, len(height)-1
	area := 0

	for start < end {
		if min(height[start], height[end])*step > area {
			area = min(height[start], height[end]) * step
		}

		// двигаем "невыгодный" указатель (тот, что меньше)
		if height[start] >= height[end] {
			end--
		} else {
			start++
		}
		step--

	}
	return area
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))       // 49
	fmt.Println(maxArea([]int{1, 1}))                            // 1
	fmt.Println(maxArea([]int{1, 2, 3, 1000, 1000, 4, 3, 2, 1})) // 1000
	fmt.Println(maxArea([]int{}))                                // 0
	fmt.Println(maxArea([]int{-3, -4}))                          // 0
	fmt.Println(maxArea([]int{1, 3, 2, 5, 25, 24, 5}))           // 24
}
