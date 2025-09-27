package main

import (
	"fmt"
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	gridStrings := make([][]string, len(grid))
	rowMap := make(map[string]int)
	colMap := make(map[string]int)

	for i := range grid {
		gridStrings[i] = make([]string, len(grid[i]))
		for j := range grid[i] {
			gridStrings[i][j] = strconv.Itoa(grid[i][j])
		}
	}
	count := 0
	for i := 0; i < len(gridStrings); i++ {
		str := strings.Join(gridStrings[i], ", ")
		rowMap[str]++

		strs := make([]string, len(gridStrings))
		for j := 0; j < len(gridStrings[i]); j++ {
			strs[j] = gridStrings[j][i]
		}
		str = strings.Join(strs, ", ")
		colMap[str]++

	}

	for key, value := range rowMap {
		count += value * colMap[key]
	}
	return count
}

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))                        // 1
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})) // 3
	fmt.Println(equalPairs([][]int{{11, 1}, {1, 11}}))                                       // 2
	fmt.Println(equalPairs([][]int{{13, 13}, {13, 13}}))                                     // 4
}
