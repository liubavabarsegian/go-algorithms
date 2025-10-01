package main

import "fmt"

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// минимальная сумма подмассива алгоритмом Кадана.
func minSubarraySum(nums []int) int {
	minSum := nums[0]
	currentMinSum := minSum

	for i := 1; i < len(nums); i++ {
		// либо прибавляем к текущей сумме, либо сбрасываем суммку и берем текущий элемент
		currentMinSum = min(currentMinSum+nums[i], nums[i])
		minSum = min(minSum, currentMinSum)
	}
	return minSum
}

func maxSubarraySum(nums []int) int {
	maxSum := nums[0]
	currentMaxSum := maxSum

	for i := 1; i < len(nums); i++ {
		currentMaxSum = max(currentMaxSum+nums[i], nums[i])
		maxSum = max(maxSum, currentMaxSum)
	}
	return maxSum
}

func maxSubarraySumCircular(nums []int) int {
	// если сумма всех элементов равна минимальной сумме подмассива, то все элементы отрицательные
	// тогда возвращаем максимальный элемент
	if sum(nums) == minSubarraySum(nums) {
		return maxSubarraySum(nums)
	}
	// иначе берем максимальную из двух сумм:
	// 1. максимальная сумма линейного подмассива
	// 2. максимальная суммая цикличного массива: это сумма всех элементов - сумма минимального подмассива
	return max(maxSubarraySum(nums), sum(nums)-minSubarraySum(nums))
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2})) // 3
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))     // 10
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))   // -2
	fmt.Println(maxSubarraySumCircular([]int{3, -1, 2, -1})) // 4
}
