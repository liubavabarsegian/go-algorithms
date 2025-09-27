package main

import (
	"fmt"
)

// c -> write to stack
// * -> delete from stack

func removeStars(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack = append(stack, s[i])
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeStars("leet**cod*e")) // "lecoe"
	fmt.Println(removeStars("erase*****"))  // ""
}
