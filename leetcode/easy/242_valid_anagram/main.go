package main

import "slices"

func isAnagram(s string, t string) bool {
	bytesS := []byte(s)
	bytesT := []byte(t)

	slices.Sort(bytesS)
	slices.Sort(bytesT)

	return string(bytesS) == string(bytesT)
}
