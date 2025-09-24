package main

import "fmt"

func pivotIndex(nums []int) int {
	res := make([]int, len(nums)+2)
	res[0] = 0
	res[len(res)-1] = 0

	// создаем массив с префиксными суммами
	for i := len(nums) - 1; i >= 0; i-- {
		res[i+1] = res[i+2] + nums[i]
	}

	// слева направо ищем совпадающую сумму
	for i := 0; i < len(res)-2; i++ {
		if res[i] == res[i+2] {
			return i
		}
		res[i+1] = res[i] + nums[i]
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))    // 3
	fmt.Println(pivotIndex([]int{1, 2, 3}))             // -1
	fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1})) // 0
	fmt.Println(pivotIndex([]int{-1, -1, 0, -1, -1}))   // 2
	fmt.Println(pivotIndex([]int{0, 0}))                // 0
	fmt.Println(pivotIndex([]int{-1, -1, 0, 1, 1, 0}))  // 5
}
