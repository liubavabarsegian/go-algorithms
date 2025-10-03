package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	words := strings.Fields(str)

	pairsMap := make(map[string]int)

	for _, word := range words {
		for i := 0; i < len(word)-1; i++ {
			substr := word[i : i+2]

			pairsMap[substr]++
		}
	}

	maxCount := 0
	var maxPair string
	for k, v := range pairsMap {
		if v > maxCount || v == maxCount && k > maxPair {
			maxCount = v
			maxPair = k
		}
	}

	fmt.Println(maxPair)
}
