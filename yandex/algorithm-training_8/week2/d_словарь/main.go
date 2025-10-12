package main

import (
	"fmt"
	"strings"
)

func splitText(text string, dictionary map[string]bool) string {
	dp := make([]bool, len(text)+1)
	parents := make([]int, len(text)+1)
	dp[0] = true

	for i := 0; i < len(text); i++ {
		for j := 0; j < i+1; j++ {
			if dp[j] && dictionary[text[j:i+1]] {
				dp[i+1] = true
				parents[i+1] = j

				break

			}
		}
	}

	res := make([]string, 0)
	i := len(dp) - 1
	for i > 0 {
		start := parents[i]
		res = append(res, text[start:i])
		i = start
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	// fmt.Println(dp, parents, res)
	return strings.Join(res, " ")
}

func main() {
	var text string
	var n int

	fmt.Scanf("%s\n", &text)
	fmt.Scanf("%d\n", &n)

	dictionary := make(map[string]bool)
	for i := 0; i < n; i++ {
		var word string
		fmt.Scanf("%s\n", &word)
		dictionary[word] = true
	}

	fmt.Println(splitText(text, dictionary))
}
