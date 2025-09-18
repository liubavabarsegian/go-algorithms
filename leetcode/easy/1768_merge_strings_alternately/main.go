package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	size := len(word1) + len(word2)
	minSize := min(len(word1), len(word2))
	chars := make([]byte, size)

	resultIndex := 0
	for i := range minSize {
		chars[resultIndex] = word1[i]
		chars[resultIndex+1] = word2[i]
		resultIndex += 2
	}

	if len(word1) == len(word2) {
		return string(chars)
	}

	var remainingWord string
	if len(word1) > len(word2) {
		remainingWord = word1[minSize:]
	} else {
		remainingWord = word2[minSize:]
	}

	for i := range remainingWord {
		chars[resultIndex] = remainingWord[i]
		resultIndex += 1
	}

	return string(chars)
}

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
}
