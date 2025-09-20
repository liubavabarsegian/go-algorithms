package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}

	iter, i := 1, 1
	var canCompare = true
	for i < len(chars) {
		count := 1
		for canCompare && i < len(chars) && chars[iter-1] == chars[i] {
			count++
			i++
		}
		canCompare = true
		if count > 1 {
			num := strconv.Itoa(count)
			for j := range len(num) {
				chars[iter] = num[j]
				iter++
			}
			canCompare = false
		} else {
			chars[iter] = chars[i]
			iter++
			i++
		}
	}
	return iter
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars)) // 6
	fmt.Println(chars)           //['a','2','b','2','c','3']

	chars = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(chars)) // 4
	fmt.Println(chars)           //['a','b','1','2']

	chars = []byte{'#', '#', '#'}
	fmt.Println(compress(chars)) // 2
	fmt.Println(chars)           //['#','3']

	chars = []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}
	fmt.Println(compress(chars)) // 6
	fmt.Println(chars)           //['a','3','b','2','a','2']

	chars = []byte{'e', 'e', '2', '2', '2', '2', '2', 'a', 'a'}
	fmt.Println(compress(chars)) // 6
	fmt.Println(chars)           //['e','2','2','5','a','2']

}
