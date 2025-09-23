package main

import (
	"fmt"
	"slices"
)

func maxVowels(s string, k int) int {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	count := 0
	for i := 0; i < len(s) && i < k; i++ {
		if slices.Contains(vowels, s[i]) {
			count++
		}
	}
	maxCount := count
	for i := k; i < len(s); i++ {
		if slices.Contains(vowels, s[i-k]) {
			count--
		}
		if slices.Contains(vowels, s[i]) {
			count++
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3)) // 3 ("iii")
	fmt.Println(maxVowels("aeiou", 2))     // 2
	fmt.Println(maxVowels("leetcode", 3))  //2
	fmt.Println(maxVowels("", 3))          //0
	fmt.Println(maxVowels("tryhard", 4))   //1
}
