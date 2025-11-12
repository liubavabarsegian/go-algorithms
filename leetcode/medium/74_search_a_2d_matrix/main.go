package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	// поиск по первому элементу строк
	// если больше [i][0], но меньше [i+1][0], то можно рассматривать строку
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	upperRow := 0
	lowerRow := len(matrix) - 1

	for upperRow <= lowerRow {
		mediumRow := (upperRow + lowerRow) / 2

		// fmt.Println(mediumRow, matrix[mediumRow][0] < target && (mediumRow == len(matrix)-1 || matrix[mediumRow+1][0] > target))
		if matrix[mediumRow][0] == target {
			return true
		} else if matrix[mediumRow][0] > target {
			lowerRow = mediumRow - 1
		} else if matrix[mediumRow][0] < target {
			upperRow = mediumRow + 1
			if mediumRow == len(matrix)-1 || matrix[mediumRow+1][0] > target {
				leftColumn := 0
				rightColumn := len(matrix[0]) - 1

				for leftColumn <= rightColumn {
					mediumColum := (leftColumn + rightColumn) / 2

					if matrix[mediumRow][mediumColum] == target {
						return true
					} else if matrix[mediumRow][mediumColum] > target {
						rightColumn = mediumColum - 1
					} else if matrix[mediumRow][mediumColum] < target {
						leftColumn = mediumColum + 1
					} else {
						return false
					}
				}
				return false
			}

		} else {
			return false
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))  // true
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)) // false
	fmt.Println(searchMatrix([][]int{{1}}, 2))                                               // false
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))                                            // true
	fmt.Println(searchMatrix([][]int{{1}, {3}}, 3))                                          // true
}
