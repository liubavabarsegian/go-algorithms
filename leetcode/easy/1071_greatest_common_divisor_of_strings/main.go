package main

import "fmt"

func divides(t string, s string) bool {
	iterSize := len(t)
	if len(t) == 0 || len(s) == 0 || len(s)%iterSize != 0 {
		return false
	}

	for i := 0; i+iterSize <= len(s); i += iterSize {
		if s[i:i+iterSize] != t {
			return false
		}
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	var shortestStr, longestStr string
	if len(str1) <= len(str2) {
		shortestStr = str1
		longestStr = str2
	} else {
		shortestStr = str2
		longestStr = str1
	}

	for i := len(shortestStr); i > 0; i-- {
		gcd := shortestStr[:i]
		if divides(gcd, longestStr) && divides(gcd, shortestStr) {
			return gcd
		}
	}
	return ""
}

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}
