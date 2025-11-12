package main

import "fmt"

func DFS(connections [][]int, visited *map[int]bool, i int, count *int) {
	pair := connections[i]
	fmt.Println(visited, pair, *count)
	if pair[0] == 0 || (pair[1] != 0 && !(*visited)[pair[1]] && (*visited)[pair[0]]) {
		pair[0], pair[1] = pair[1], pair[0]
		(*count)++
	}

	(*visited)[pair[0]] = true

	if !(*visited)[pair[1]] {
		(*visited)[pair[1]] = true
		// DFS(connections, visited, i, count)
	}
}

func minReorder(n int, connections [][]int) int {
	isConnected := make([][]bool, n)
	for i := range n {
		isConnected[i] = make([]bool, n)
	}
	// for i := range len(connections) {
	// 	isConnected[connections[i][0]][connections[i][1]] = true
	// 	isConnected[connections[i][1]][connections[i][0]] = true
	// }

	visited := make(map[int]bool, n)
	count := 0

	for i := range connections {
		DFS(connections, &visited, i, &count)
	}
	return count
}

// 0, 1 -> 1
func main() {
	fmt.Println(minReorder(5, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
	fmt.Println(minReorder(3, [][]int{{1, 0}, {2, 0}}))
	fmt.Println(minReorder(3, [][]int{{1, 0}, {2, 0}}))
}
