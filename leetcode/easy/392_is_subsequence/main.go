package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == t || len(s) == 0 {
		return true
	}

	subIter := 0
	for i := 0; i < len(t); i++ {
		if subIter < len(s) && s[subIter] == t[i] {
			subIter++
		}

		if subIter == len(s) {
			return true
		}

	}

	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
	fmt.Println(isSubsequence("ace", "abcde"))  // true
	fmt.Println(isSubsequence("aec", "abcde"))  // false
	fmt.Println(isSubsequence("", "abcde"))     // true
	fmt.Println(isSubsequence("", ""))          // true
	fmt.Println(isSubsequence("b", "c"))        // false
}
