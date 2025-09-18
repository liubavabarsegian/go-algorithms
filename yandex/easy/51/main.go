package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	words := strings.Fields(string(file))

	wordsMap := make(map[string]int)
	results := make([]int, 0)
	for _, w := range words {
		if _, ok := wordsMap[w]; ok {
			wordsMap[w] += 1
			results = append(results, wordsMap[w])
		} else {
			wordsMap[w] = 0
			results = append(results, 0)
		}
	}

	resultsString := make([]string, len(results))
	for i := range results {
		resultsString[i] = strconv.Itoa(results[i])
	}

	writer.WriteString(strings.Join([]string(resultsString), " "))
	writer.WriteByte('\n')
}
