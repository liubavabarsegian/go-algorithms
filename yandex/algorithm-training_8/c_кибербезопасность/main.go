package main

import "fmt"

func passwordsCount(str string) uint64 {
	chars := make(map[byte]uint64)

	for i := 0; i < len(str); i++ {
		chars[str[i]]++
	}

	totalCombinations := uint64(len(str) * (len(str) - 1) / 2)
	for _, v := range chars {
		if v > 0 {
			uselessCombinations := v * (v - 1) / 2

			totalCombinations -= uselessCombinations
		}

	}
	return totalCombinations + 1
}

func main() {
	var str string
	fmt.Scanf("%s\n", &str)
	fmt.Println(passwordsCount(str))
}
