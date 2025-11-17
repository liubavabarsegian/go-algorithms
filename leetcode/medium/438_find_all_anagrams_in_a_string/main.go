package main

import (
	"fmt"
	"slices"
)

func findAnagrams1(s string, p string) []int {
	anagramSize := len(p)
	if anagramSize == 0 {
		return []int{}
	}

	bytesP := []byte(p)
	slices.Sort(bytesP)

	res := []int{}
	for i := 0; i < len(s)-anagramSize+1; i++ {
		bytesSubStr := []byte(s[i : i+anagramSize])
		slices.Sort(bytesSubStr)

		if string(bytesSubStr) == string(bytesP) {
			res = append(res, i)
		}
	}
	return res
}

func findAnagrams(s string, p string) []int {
	freqS := make([]int, 26)
	freqP := make([]int, 26)

	if len(s) < len(p) || len(p) == 0 {
		return []int{}
	}

	for i := 0; i < len(p); i++ {
		freqS[int(s[i]-'a')]++
		freqP[int(p[i]-'a')]++
	}

	start := 0
	end := len(p)

	res := []int{}
	if fmt.Sprint(freqS) == fmt.Sprint(freqP) {
		res = append(res, start)
	}

	for end < len(s) {
		freqS[int(s[start]-'a')]--
		freqS[int(s[end]-'a')]++

		if fmt.Sprint(freqS) == fmt.Sprint(freqP) {
			res = append(res, start+1)
		}

		start++
		end++
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams("abab", "ab"))        //[0,1,2]
	fmt.Println(findAnagrams("", "ab"))            //[]
	fmt.Println(findAnagrams("abc", ""))           //[]
}
