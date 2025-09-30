package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(str string) {
	s.items = append(s.items, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	top := s.Top()
	s.items = s.items[0 : len(s.items)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Top() string {
	return s.items[len(s.items)-1]
}

// num -> положить в стек
// [ -> положить в стек
// char -> положить в стек
// ] -> заворачиваем

func extractNumber(s string) string {
	i := 0
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		i++
	}
	return s[0:i]
}

func extractChars(s string) string {
	i := 0
	for i < len(s) {
		if s[i] < 'a' || s[i] > 'z' {
			break
		}
		i++
	}
	return s[:i]
}

func isAlphaNumeric(s string) bool {
	for _, v := range s {
		if v < 'a' || v > 'z' {
			return false
		}
	}
	return true
}

func decodeString(s string) string {
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := extractNumber(s[i:])
			stack.Push(num)
			i += len(num) - 1
		}
		if s[i] == '[' {
			stack.Push(string(s[i]))
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			chars := extractChars(s[i:])
			stack.Push(chars)
			i += len(chars) - 1
		}
		for s[i] == ']' && len(stack.items) > 2 && isAlphaNumeric(stack.Top()) {
			// начинаем схлопывать строку
			// последний элемент стека - символы, достаем их
			chars := stack.Pop()

			// если есть что склеить слева - склеиваем
			if !stack.IsEmpty() && isAlphaNumeric(stack.Top()) {
				chars = stack.Pop() + chars
			}
			// если можем их схлопнуть, то начинаем
			if stack.Top() == "[" {
				stack.Pop()
				//  достаем последние символы и умножаем на число
				num := stack.Pop()
				numInt, _ := strconv.Atoi(num)
				chars = strings.Repeat(chars, numInt)

				stack.Push(chars)
				break
			}
			stack.Push(chars)
		}
	}
	return strings.Join(stack.items, "")
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))                    // "aaabcbc"
	fmt.Println(decodeString("3[a2[c]]"))                     // "accaccacc"
	fmt.Println(decodeString("2[abc]3[cd]ef"))                // "abcabccdcdcdef"
	fmt.Println(decodeString("2[a2[b2[c]]]"))                 // "abccbccabccbcc"
	fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")) // "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"
}
