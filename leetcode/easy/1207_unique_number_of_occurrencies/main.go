package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	occurences := make(map[int]uint)
	for _, v := range arr {
		occurences[v]++
	}

	occurencesOfOccurences := make(map[uint]bool)
	for _, v := range occurences {
		if occurencesOfOccurences[v] {
			return false
		}
		occurencesOfOccurences[v] = true
	}
	return true
}

func main() {
	// The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})) // true.
	fmt.Println(uniqueOccurrences([]int{1, 2}))             // false
}
