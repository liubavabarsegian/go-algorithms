package main

import (
	"fmt"
)

func bestPath(a, b, c, v0, v1, v2 float64) float64 {
	minSum := a/v0 + c/v1 + b/v2
	minSum = min(minSum, b/v0+c/v1+a/v2)
	minSum = min(minSum, a/v0+a/v1+b/v0+b/v1)
	minSum = min(minSum, a/v0+c/v0+b/v1+a/v0+a/v1)
	minSum = min(minSum, b/v0+c/v0+a/v1+b/v0+b/v1)
	minSum = min(minSum, b/v0+c/v1+c/v2+b/v2)
	minSum = min(minSum, a/v0+c/v1+c/v2+a/v2)

	return minSum
}

func main() {
	var a, b, c, v0, v1, v2 float64
	fmt.Scanf("%f %f %f %f %f %f\n", &a, &b, &c, &v0, &v1, &v2)
	fmt.Println(bestPath(a, b, c, v0, v1, v2))
}
