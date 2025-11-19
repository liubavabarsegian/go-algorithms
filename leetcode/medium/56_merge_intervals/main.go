package main

import (
	"fmt"
	"sort"
)

// нужно отсортировать по 1 элементу
// потом объединять
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := make([][]int, 0, 1)
	res = append(res, intervals[0])
	resIter := 0

	for i := 1; i < len(intervals); i++ {
		if resIter < len(res) && res[resIter][1] >= intervals[i][0] {
			res[resIter][1] = max(res[resIter][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			resIter++
		}
	}
	return res
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1,5]]
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))                    // [[1,7]]
	fmt.Println(merge([][]int{{4, 7}, {4, 5}}))                    // [[4,7]]
	fmt.Println(merge([][]int{{4, 7}}))                            // [[4,7]]
	fmt.Println(merge([][]int{{}}))                                // [[]]
	fmt.Println(merge([][]int{{1, 4}, {0, 5}}))                    // [[0, 5]]
}
