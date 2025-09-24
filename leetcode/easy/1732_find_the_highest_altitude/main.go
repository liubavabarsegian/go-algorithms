package main

import "fmt"

func largestAltitude(gain []int) int {
	largestAltitude := 0
	prefixSum := 0
	for i := 0; i < len(gain); i++ {
		prefixSum += gain[i]
		largestAltitude = max(largestAltitude, prefixSum)
	}

	return largestAltitude
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))         // 1
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2})) // 0

}
