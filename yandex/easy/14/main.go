package main

import (
	"fmt"
)

type Point struct {
	row    int
	column int
}

/*
	a . . .
	. . f .
	b . . .
*/

// count moves from feedPoint to point
func countMoves(point Point, feedPoint Point, rows, columns int) {
	// base case

	// the rest logic
}

func main() {
	var rows, columns, count int
	feedPoint := Point{}
	fleasPoints := make([]Point, 0)

	fmt.Scanf("%d %d %d %d %d'n", &rows, &columns, &feedPoint.row, &feedPoint.column, &count)
	for range count {
		point := Point{}
		fmt.Scanf("%d %d\n", &point.row, &point.column)
		fleasPoints = append(fleasPoints, point)
	}
}
