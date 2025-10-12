package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	n_input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(n_input))

	mushrooms_input, _ := reader.ReadString('\n')
	mushrooms_strings := strings.Split(strings.TrimSpace(mushrooms_input), " ")

	mushrooms := make([]int, len(mushrooms_strings))
	for i := 0; i < len(mushrooms_strings); i++ {
		mushroom, _ := strconv.Atoi(mushrooms_strings[i])
		mushrooms[i] = mushroom
	}

	writer.WriteString(strconv.Itoa(countVasyasHappiness(n, mushrooms)))
	writer.WriteByte('\n')
}
