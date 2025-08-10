package main

import "fmt"

func rotate(nums []int, k int) {

	for k > len(nums) {
		k -= len(nums)
	}
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)

}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		swap(&nums[i], &nums[n-i-1])
	}
}

func swap(n1 *int, n2 *int) {
	*n1, *n2 = *n2, *n1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{-1}
	rotate(nums, 2)
	fmt.Println(nums)

	nums = []int{1, 2}
	rotate(nums, 3)
	fmt.Println(nums)
}
