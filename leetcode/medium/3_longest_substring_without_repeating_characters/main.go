package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start := 0
	end := 1

	uniqueMap := make(map[byte]int)
	uniqueMap[s[0]] = 0
	maxLen := 1
	for end < len(s) {
		if _, ok := uniqueMap[s[end]]; !ok {
			uniqueMap[s[end]] = end
			end++
		} else {
			if uniqueMap[s[end]] >= start {
				start = uniqueMap[s[end]] + 1
			}

			uniqueMap[s[end]] = end
			end++
		}
		maxLen = max(maxLen, end-start)
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring("pwewkwew")) // 3
	fmt.Println(lengthOfLongestSubstring(""))         // 0
	fmt.Println(lengthOfLongestSubstring("v"))        // 1
}
