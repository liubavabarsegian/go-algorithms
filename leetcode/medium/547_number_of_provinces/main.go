package main

import "fmt"

func DFS(isConnected [][]int, visited *[]bool, city int) {
	(*visited)[city] = true

	for i, v := range isConnected[city] {
		if i != city && !(*visited)[i] && v == 1 {
			DFS(isConnected, visited, i)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	provinces := 0

	for i := range isConnected {
		if !visited[i] {
			DFS(isConnected, &visited, i)
			provinces++
		}

	}
	return provinces
}

/*
	1 1 0
	1 1 0
	0 0 1
*/

/*
	1 0 0
	0 1 0
	0 0 1
*/

/*
	1 0 0 1
	0 1 1 0
	0 1 1 1
	1 0 1 1
*/

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))                        // 2
	fmt.Println(findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))                        // 3
	fmt.Println(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}})) // 1
}
