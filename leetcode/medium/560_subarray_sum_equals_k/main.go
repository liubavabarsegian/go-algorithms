package main

import "fmt"

func subarraySum(nums []int, k int) int {
	frequenciesMap := make(map[int]int)

	prefixSum, result := 0, 0
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		// we update the result only when prefixSum - k exists in the map
		if _, ok := frequenciesMap[prefixSum-k]; ok {
			result += frequenciesMap[prefixSum-k]
		}
		if prefixSum == k {
			result++
		}
		frequenciesMap[prefixSum]++
	}
	return result
}

/*
-1 -> 1
0 -> 1 res = 1
0 -> 2 res = 2
*/
func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))              // 2
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))              // 2
	fmt.Println(subarraySum([]int{1, 2, 3, 1, 2, 3}, 3))     // 4
	fmt.Println(subarraySum([]int{10, 2, -2, -20, 10}, -10)) // 3
	fmt.Println(subarraySum([]int{-1, 1, 0}, 0))             // 3
}
