package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	brickSides := make([]int, 0, 3)
	holeSides := make([]int, 0, 2)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	for range 3 {
		input, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(input))
		brickSides = append(brickSides, num)
	}

	for range 2 {
		input, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(input))
		holeSides = append(holeSides, num)
	}

	slices.Sort(brickSides)
	slices.Sort(holeSides)

	okayCount := 0
	for _, v := range holeSides {
		for i, k := range brickSides {
			if k <= v {
				okayCount += 1
				brickSides = slices.Delete(brickSides, i, i+1)
				break
			}
		}
	}

	if okayCount == 2 {
		writer.WriteString("YES")
	} else {
		writer.WriteString("NO")
	}

	writer.WriteByte('\n')
}
