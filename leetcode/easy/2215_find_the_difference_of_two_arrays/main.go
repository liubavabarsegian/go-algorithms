package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	res1 := make([]int, 0, len(nums1))
	res2 := make([]int, 0, len(nums2))

	for _, v := range nums1 {
		map1[v] = true
	}
	for _, v := range nums2 {
		if !map2[v] {
			map2[v] = true
			if !map1[v] {
				res2 = append(res2, v)
			}

		}
	}
	for key := range map1 {
		if !map2[key] {
			res1 = append(res1, key)
		}
	}

	return [][]int{res1, res2}

}

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))       // [[1,3],[4,6]]
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})) // [[3],[]]
	fmt.Println(findDifference([]int{}, []int{1, 1, 2, 2}))           // [[0],[1, 2]]
}
