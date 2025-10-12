package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// to cover as more topics as possible
func pickTasks(k int, tasks []uint64) []uint64 {
	tasksMap := make(map[uint64]uint64)
	for i := 0; i < len(tasks); i++ {
		tasksMap[tasks[i]]++
	}

	results := make([]uint64, k)
	resultsCount := 0
	for key := range tasksMap {
		if resultsCount >= k {
			break
		}

		results[resultsCount] = key
		tasksMap[key]--
		resultsCount++
	}
	for key := range tasksMap {
		for resultsCount < k && tasksMap[key] > 0 {
			results[resultsCount] = key
			resultsCount++
			tasksMap[key]--
		}
	}

	return results
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n, k int

	fmt.Scanf("%d %d\n", &n, &k)
	tasks_input, _ := reader.ReadString('\n')

	tasks_strs := strings.Split(strings.TrimSpace(tasks_input), " ")
	tasks := make([]uint64, len(tasks_strs))
	for i := 0; i < len(tasks_strs); i++ {
		num, _ := strconv.Atoi(tasks_strs[i])
		tasks[i] = uint64(num)
	}

	result := pickTasks(k, tasks)
	out := make([]string, len(result))
	for i, v := range result {
		out[i] = strconv.FormatUint(v, 10)
	}
	fmt.Fprintln(writer, strings.Join(out, " "))
}
