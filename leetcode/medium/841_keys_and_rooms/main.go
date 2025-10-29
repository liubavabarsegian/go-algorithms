package main

import (
	"fmt"
	"slices"
)

func DFS(rooms [][]int, visited *[]bool, room int) {
	(*visited)[room] = true

	for _, v := range rooms[room] {
		if !(*visited)[v] {
			DFS(rooms, visited, v)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	DFS(rooms, &visited, 0)

	return !slices.Contains(visited, false)
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
