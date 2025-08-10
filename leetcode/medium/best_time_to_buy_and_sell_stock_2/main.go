package main

import "fmt"

// [7, 1, 5, 3, 6, 4]
// 7: min = 7 prev = 7 sumProfit = 0
// 1: min = 1 prev = 7 sumProfit = 0
// 5: min = 1 prev = 1 sumProfit = 0
// 3: min = 1 prev = 5 sumProfit = 4
// 6: min = 3 prev = 3 sumProfit = 4
// 4: min = 3 prev = 6 sumProfit = 7
// ---------------------------------
// [2, 4, 1]
// 2: min = 2 prev = 2 sumProfit = 0
// 4: min = 2 prev = 2 sumProfit = 0
// 1: min = 2 prev = 4 sumProfit = 2
// ---------------------------------
// [1, 2, 3, 4, 5]
// 1: min = 1 prev = 1 sumProfit = 0
// 2: min = 1 prev = 1 sumProfit = 0
// 3: min = 1 prev = 2 sumProfit = 0
// 4: min = 1 prev = 2 sumProfit = 0
// 5: min = 1 prev = 2 sumProfit = 0
func maxProfit(prices []int) int {
	sumProfit := 0
	minPrice := prices[0]
	prevPrice := prices[0]

	for day, price := range prices {
		if day > 1 && price < prevPrice {
			sumProfit += prevPrice - minPrice
			minPrice = price

		} else if day == len(prices)-1 && price >= prevPrice {
			sumProfit += price - minPrice
		}

		if price < minPrice {
			minPrice = price
		}
		prevPrice = price

	}
	return sumProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

	prices = []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}
