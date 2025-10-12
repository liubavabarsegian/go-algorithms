package main

import "fmt"

// LLBLRRBRL - 5
// 	[L, L, B, L, R, R, B, R, L]
/*
	[1, 1, 2, 2, 3, 3, 4, 4, 5]
*/

// LLBLRRBRLR - 6
// 	[L, L, B, L, R, R, B, R, L, R]
/*
	[1, 1, 2, 2, 3, 3, 4, 4, 5, R]
*/

// LLBLRRBRLRR - 6 или 7
// 	[L, L, B, L, R, R, B, R, L, R]
/*
	[1, 1, 2, 2, 3, 3, 4, 4, 5, R]
*/

// LRLRLRLRLRLRLR - 8
// 	[L, R, L, R, L, R, L, R, L, R, L, R, L, R]
/*
	[1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 14]
*/

// BRRLRLBBLB - 6
// 	[B, R, R, L, R, L, B, B, L, B]
/*
	[1, 1, 1, 2, 2, 3, 4, 5, 5, 6]
*/
func countCrossings(inflows string) int {
	// это массив для минимального количества переправ, которое нужно для того, чтобы оказаться в точке i
	// [i][j]
	// j - это для левого или правого берега
	minCrossings := make([][2]int, len(inflows)+1)
	minCrossings[0][0] = 0
	minCrossings[0][1] = 1
	for i := 0; i < len(inflows); i++ {
		switch inflows[i] {
		case 'B':
			minCrossings[i+1][0] = min(minCrossings[i][0]+1, minCrossings[i][1]+2)
			minCrossings[i+1][1] = min(minCrossings[i][0]+2, minCrossings[i][1]+1)
		case 'R':
			minCrossings[i+1][0] = min(minCrossings[i][0], minCrossings[i][1]+1)
			minCrossings[i+1][1] = min(minCrossings[i][0]+1, minCrossings[i][1]+1)
		case 'L':
			minCrossings[i+1][0] = min(minCrossings[i][0]+1, minCrossings[i][1]+1)
			minCrossings[i+1][1] = min(minCrossings[i][0]+1, minCrossings[i][1])
		}
	}

	return minCrossings[len(minCrossings)-1][1]
}

func main() {
	var inflows string

	fmt.Scanf("%s\n", &inflows)
	fmt.Println(countCrossings(inflows))
}
