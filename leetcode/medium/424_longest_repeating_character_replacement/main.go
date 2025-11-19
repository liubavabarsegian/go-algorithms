package main

import (
	"fmt"
)

// sliding window
// если использование k == 0, то сдвигаем левый указатель
func maxFrequency(s string) int {
	freq := make(map[byte]int)
	maxFreq := 0
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
		if freq[s[i]] > maxFreq {
			maxFreq = freq[s[i]]
		}
	}
	return maxFreq
}

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	freq := make(map[byte]int, 26)
	start, end := 0, 0
	maxLen, maxFreq := 0, 0

	for end < len(s) {
		freq[s[end]]++
		maxFreq = max(maxFreq, freq[s[end]])

		// invalid window
		// fmt.Println(end-start+1-maxFreq > k, end-start+1, maxFreq, s[start:end+1], start, end)

		for end-start+1-maxFreq > k {
			freq[s[start]]--
			start++

			maxFreq = max(maxFreq, freq[s[start]])
		}

		maxLen = max(maxLen, end-start+1)
		end++

	}
	return maxLen
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))      // 4
	fmt.Println(characterReplacement("AABABBA", 1))   // 4
	fmt.Println(characterReplacement("AAACBABBA", 1)) // 4
	fmt.Println(characterReplacement("", 1))          // 0
	fmt.Println(characterReplacement("EER", 0))       // 2
	fmt.Println(characterReplacement("EEEE", 0))      // 4
}
