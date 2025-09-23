package main

import "fmt"

func largestAltitude(gain []int) int {
	res := make([]int, len(gain)+1)

	res[0] = 0
	largestAltitude := 0
	for i := 0; i < len(gain); i++ {
		res[i+1] = res[i] + gain[i]

		largestAltitude = max(largestAltitude, res[i+1])
	}

	return largestAltitude
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))         // 1
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2})) // 0

}
