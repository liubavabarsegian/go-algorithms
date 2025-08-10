package main

import (
	"fmt"
	"slices"
)

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}

func isVowel(c byte) bool {
	return slices.Contains(vowels, c) || slices.Contains(vowels, c+32)
}

func reverseVowels(s string) string {

	start := 0
	end := len(s) - 1
	result := make([]byte, len(s))

	for i := range len(s) {
		result[i] = s[i]
	}

	for {
		if start >= end {
			break
		}

		for !isVowel(result[start]) && start < end {
			start++
		}

		for !isVowel(result[end]) && start < end {
			end--
		}

		if isVowel(result[start]) && isVowel(result[end]) {
			result[start], result[end] = result[end], result[start]
			start++
			end--
		}
	}

	return string(result)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("lttt"))
}
