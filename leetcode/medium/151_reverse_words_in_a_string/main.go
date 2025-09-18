package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	results := make([]string, len(words))

	for i := range len(words) {
		results[i] = words[len(words)-1-i]
	}

	return strings.Join(results, " ")
}

func main() {
	fmt.Println(reverseWords(" hello 	world   "))
}
