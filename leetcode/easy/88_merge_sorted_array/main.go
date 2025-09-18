package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	commonIter := m + n - 1
	nums1Iter := m - 1
	nums2Iter := n - 1
	i, j := nums1Iter, nums2Iter

	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[commonIter] = nums1[i]
			i--

		} else {

			nums1[commonIter] = nums2[j]
			j--
		}
		commonIter--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)

	nums1 = []int{2, 0}
	nums2 = []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
