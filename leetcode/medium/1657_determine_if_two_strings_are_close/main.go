package main

import (
	"fmt"
	"slices"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	map1 := make(map[byte]int)
	map2 := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		map1[word1[i]]++
		map2[word2[i]]++
	}

	counts1 := make([]int, 0)
	for i, v := range map1 {
		if map2[i] == 0 {
			return false
		}
		counts1 = append(counts1, v)
	}

	counts2 := make([]int, 0)
	for i, v := range map2 {
		if map1[i] == 0 {
			return false
		}
		counts2 = append(counts2, v)
	}

	slices.Sort(counts1)
	slices.Sort(counts2)

	return slices.Equal(counts1, counts2)
}

func main() {
	fmt.Println(closeStrings("abc", "bca"))         // true
	fmt.Println(closeStrings("a", "aa"))            // false
	fmt.Println(closeStrings("cabbba", "abbccc"))   // true
	fmt.Println(closeStrings("abbzzca", "babzzcz")) // false
}

// a bb zz c a
// b a b zz c z
