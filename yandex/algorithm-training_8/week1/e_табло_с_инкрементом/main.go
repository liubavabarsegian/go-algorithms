package main

import "fmt"

func countSeconds(number, seconds uint64) uint64 {
	for range min(3, seconds) {
		number += number % 10
	}

	if seconds <= 3 || number%10 == 0 {
		return number
	}

	if number%10 == 5 || number%10 == 10 {
		return 10
	}

	seconds -= 3
	for seconds > 5 {
		number += 20
		seconds -= 4
	}

	for range seconds {
		number += number % 10
	}

	return number
}

func main() {
	var n, k uint64

	fmt.Scanf("%d %d\n", &n, &k)
	fmt.Println(countSeconds(n, k))
}

// 1: 1, 2, 4, 8, 16, 22, 24, 28, 36, 42...
// 2 (как 1): 2, 4, 8, и тд
// 3: 3, 6, 12, 14, 18, 26, 32, 34, 38, 46
// 4(как 1): 4, 8, 16, и тд
// 5: 5, 10, 10 , 10
// 6(как 3): 6, 12, 14, 18 и тд
// 7(как 3): 7, 14, 18, 26 и тд
// 8(как 1): 8, 16, 22, 24 и тд
// 9(как 3): 9, 18, 26, 32 и тд
// 10: 10, 10, 10
