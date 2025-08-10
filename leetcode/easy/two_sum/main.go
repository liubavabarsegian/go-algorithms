package main

func twoSumSorted(nums []int, target int) []int {
	var result []int

	left := 0
	right := len(nums) - 1

	for left < right {
		switch {
		case nums[left]+nums[right] == target:
			return append(result, left, right)
		case nums[left]+nums[right] < target:
			left++
		case nums[left]+nums[right] > target:
			right--
		}
	}
	return result
}

func twoSumUnsorted(nums []int, target int) []int {
	matches := make(map[int]int, len(nums))

	for index, value := range nums {
		if matchedIndex, ok := matches[value]; ok {
			return []int{matchedIndex, index}
		}
		matches[target-value] = index
	}

	return nil
}
