package main

import "fmt"

// my first solution
func maxProfit1(prices []int) int {
	var maxProfit int
	for i := 0; i < len(prices)-1; i++ {
		profit := findMax(prices[i+1:]) - prices[i]
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func findMax(prices []int) int {
	max := prices[0]

	for _, value := range prices {
		if value > max {
			max = value
		}
	}
	return max
}

// optimized solution
// [2 4 1]
// minPrice = 2 maxProfit = 0
// minPrice = 2 maxProfit = 2
// minPirce = 1 maxProfit = 2
// --------------------------
// [7 1 5 3 6 4]
// minPrice = 7 maxProfit = 0
// minPrice = 1 maxProfit = 0
// minPrice = 1 maxProfit = 4
// minPrice = 1 maxProfit = 4
// minPrice = 1 maxProfit = 5
// minPrice = 1 maxProfit = 5
func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

	prices = []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}
