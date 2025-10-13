package main

import (
	"fmt"
	"strconv"
)

// start - end -> один символ
// end - start - range

func summaryRanges(nums []int) []string {
	ranges := make([]string, 0, len(nums))

	start, end := 0, 1
	for end <= len(nums) {
		if end < len(nums) && nums[end]-nums[end-1] == 1 {
			end++
		}

		if end == len(nums) || nums[end]-nums[end-1] > 1 {
			var str string

			if len(nums[start:end]) > 1 {
				str = fmt.Sprintf("%d->%d", nums[start], nums[end-1])
			}
			if len(nums[start:end]) == 1 {
				str = strconv.Itoa(nums[start])
			}
			ranges = append(ranges, str)
			start = end
			end++
		}
	}

	return ranges
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))    // ["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // ["0","2->4","6","8->9"]
	fmt.Println(summaryRanges([]int{1, 2, 3, 4, 5}))       // ["1->5"]
	fmt.Println(summaryRanges([]int{0}))                   // ["0"]
	fmt.Println(summaryRanges([]int{}))                    // []
}
