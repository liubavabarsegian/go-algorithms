package main

import (
	"fmt"
)

func removeDuplicateLetters(s string) string {
	hashMap := make(map[byte]int)
	stackMap := make(map[byte]int)
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if i > 0 {
			// если символ меньше последнего
			if stackMap[s[i]] == 0 && s[i] < stack[len(stack)-1] {
				// пока в стеке символы больше и имеют запас, очищаем стек
				for len(stack) > 0 && s[i] < stack[len(stack)-1] && hashMap[stack[len(stack)-1]] > 0 {
					stackMap[stack[len(stack)-1]] = 0
					stack = stack[0 : len(stack)-1]
				}
			}
		}
		if stackMap[s[i]] == 0 {
			stack = append(stack, s[i])
			stackMap[s[i]]++
		}

		hashMap[s[i]]--
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicateLetters("bcabc") == "abc")     // abc
	fmt.Println(removeDuplicateLetters("cbacdcbc") == "acdb") // acdb
	fmt.Println(removeDuplicateLetters("cdadabcc") == "adbc") // adbc
	fmt.Println(removeDuplicateLetters("abacb") == "abc")     // abc
}
