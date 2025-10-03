package main

import (
	"fmt"
	"strconv"
)

func countVasyasHappiness(n int, mushrooms []int) int {
	sum := 0
	var minVasya, maxMasha int
	if n >= 1 {
		minVasya = mushrooms[0]
	}
	if n >= 2 {
		maxMasha = mushrooms[1]
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			minVasya = min(mushrooms[i], minVasya)
			sum += mushrooms[i]
		} else {
			maxMasha = max(mushrooms[i], maxMasha)
			sum -= mushrooms[i]
		}

	}

	if maxMasha > minVasya {
		sum = sum + 2*maxMasha - 2*minVasya
	}
	return sum
}

func main() {

	fmt.Scanf("%d %d %d %d %d'n")

	writer.WriteString(strconv.Itoa(countVasyasHappiness(n, mushrooms)))
	writer.WriteByte('\n')
}
