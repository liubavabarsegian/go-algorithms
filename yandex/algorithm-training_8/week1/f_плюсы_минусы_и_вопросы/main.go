package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func countMaxRow(row []byte) int {
	count := 0
	for i := 0; i < len(row); i++ {
		if row[i] == '?' || row[i] == '+' {
			count++
		} else {
			count--
		}
	}
	return count
}

func countMinColumns(matrix [][]byte) []int {
	minColumns := make([]int, len(matrix[0]))

	for j := 0; j < len(matrix[0]); j++ {
		count := 0
		for i := 0; i < len(matrix); i++ {

			if matrix[i][j] == '-' || matrix[i][j] == '?' {
				count--
			} else {
				count++
			}
		}
		minColumns[j] = count
	}
	return minColumns
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscan(in, &m, &n)
	if m == 0 || n == 0 {
		return
	}

	matrix := make([][]byte, m)
	maxRows := make([]int, m)
	for i := 0; i < m; i++ {
		var row string
		fmt.Fscan(in, &row)
		matrix[i] = []byte(row)

		maxRows[i] = countMaxRow(matrix[i])
	}

	minColumns := countMinColumns(matrix)

	maxDiff := math.MinInt

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var contribRow, contribCol int

			switch matrix[i][j] {
			case '+':
				contribRow = 1
				contribCol = 1
			case '-':
				contribRow = -1
				contribCol = -1
			default:
				contribRow = 1
				contribCol = -1
			}

			rowWithoutJ := maxRows[i] - contribRow
			colWithoutI := minColumns[j] - contribCol

			diff := rowWithoutJ - colWithoutI
			maxDiff = max(maxDiff, diff)
		}
	}

	fmt.Println(maxDiff)
}
