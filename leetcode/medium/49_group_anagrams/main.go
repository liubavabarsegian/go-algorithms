package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, str := range strs {
		byteStr := []byte(str)
		slices.Sort(byteStr)

		if _, ok := group[string(byteStr)]; ok {
			group[string(byteStr)] = append(group[string(byteStr)], str)
		} else {
			group[string(byteStr)] = []string{str}
		}
	}

	res := [][]string{}
	for _, v := range group {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"", ""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
