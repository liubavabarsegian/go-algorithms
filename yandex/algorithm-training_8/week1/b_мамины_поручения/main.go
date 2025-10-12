package main

import (
	"fmt"
)

func bestPath(a, b, c, v0, v1, v2 float64) float64 {
	min_home_supermarket := min(a, b+c)
	min_home_pvz := min(b, a+c)
	min_supermarket_pvz := min(c, a+b)

	minSum := min_home_supermarket/v0 + min_supermarket_pvz/v1 + min_home_pvz/v2
	minSum = min(minSum, min_home_pvz/v0+min_supermarket_pvz/v1+min_home_supermarket/v2)
	minSum = min(minSum, min_home_supermarket/v0+min_home_supermarket/v1+min_home_pvz/v0+min_home_pvz/v1)

	return minSum
}

func main() {
	var a, b, c, v0, v1, v2 float64
	fmt.Scanf("%f %f %f %f %f %f\n", &a, &b, &c, &v0, &v1, &v2)
	fmt.Println(bestPath(a, b, c, v0, v1, v2))
}
