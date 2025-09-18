package main

import (
	"fmt"
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	withExtraCandies := make([]int, len(candies))
	results := make([]bool, len(candies))

	for i, c := range candies {
		withExtraCandies[i] = c + extraCandies
	}

	for i, c := range withExtraCandies {
		results[i] = slices.Max(append([]int{c}, candies...)) == c
	}
	return results
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	fmt.Println(kidsWithCandies(candies, 3))

	candies = []int{4, 2, 1, 1, 2}
	fmt.Println(kidsWithCandies(candies, 1))
}
