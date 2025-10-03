package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countCurrentMoves(row, col int, moves [][]int) int {
	counter := 0

	if row-2 >= 0 && col-1 >= 0 {
		counter += moves[row-2][col-1]
	}
	if row-1 >= 0 && col-2 >= 0 {
		counter += moves[row-1][col-2]
	}

	return counter
}

func countMoves(rows, cols int) int {
	moves := make([][]int, rows)
	for r := range rows {
		moves[r] = make([]int, cols)
	}

	moves[0][0] = 1
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			moves[row][col] = countCurrentMoves(row, col, moves)
		}
	}

	return moves[rows-1][cols-1]

}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	n_string, _ := reader.ReadString(' ')
	m_string, _ := reader.ReadString('\n')

	n, _ := strconv.Atoi(strings.TrimSpace(n_string))
	m, _ := strconv.Atoi(strings.TrimSpace(m_string))

	writer.WriteString(strconv.Itoa(countMoves(n, m)))
	writer.WriteByte('\n')

}
