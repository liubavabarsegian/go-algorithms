package main

import "fmt"

func fourSum(nums []int, target int) [][]int {

}

func main() {
	// -2 -1 0 0 1 2
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // [[2,2,2,2]]
}
