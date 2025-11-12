package main

import "fmt"

// ababcc
// [a, b]
// [aba]
// [abab]
// [abab, c]
// [abab, cc]
func partitionLabels(s string) []int {
	lastOccurences := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		lastOccurences[s[i]] = i
	}

	results := []int{}
	waitIter := 0
	partitionsIter := 0
	for i := 0; i < len(s); i++ {
		if waitIter <= lastOccurences[s[i]] {
			waitIter = lastOccurences[s[i]]
		}
		// fmt.Println(waitIter, string(s[i]), lastOccurences[s[i]])
		if waitIter <= i {
			results = append(results, i-partitionsIter+1)
			partitionsIter = i + 1
		}
	}
	return results
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij")) // [9, 7, 8]
	fmt.Println(partitionLabels("eccbbbbdec"))               // [10]
}
