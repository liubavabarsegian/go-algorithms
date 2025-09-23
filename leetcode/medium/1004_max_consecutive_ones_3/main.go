package main

import "fmt"

// проще осознать эту задачу, если сформулировать ее для себя как
// "самый длинный размер подмассива с максимальным k-количеством нулей"
// используя sliding window, мы скользим, изменяя размер окна
func longestOnes(nums []int, k int) int {
	replaced, count, maxCount := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		}
		if nums[i] == 0 {
			switch {
			// пока замен с запасом - заменяем 0 и увеличиваем счетчик
			case replaced < k:
				replaced++
				count++
			// замены закончились. этот кейс как бы можно пропустить, но оставила для наглядности
			// по идее, если мы снова на нуле и самый левый граничный элемент подмассива тоже ноль,
			// то мы не совершаем лишних действий, просто мысленно выкидываем граничный ноль и заменяем текущий
			case replaced == k && nums[i-count] == 0:
				continue
			// замены закончились. придется отбрасывать все единички до нуля и уменьшать счетчик
			case replaced == k && nums[i-count] == 1:
				for nums[i-count] == 1 {
					count--
				}
			}

		}

		maxCount = max(maxCount, count)
	}
	return maxCount
}

// 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))                         // 6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) // 10
	fmt.Println(longestOnes([]int{1, 1, 0, 1, 1, 1}, 0))                                        // 3
	fmt.Println(longestOnes([]int{1, 0, 1, 1, 0, 1}, 0))                                        // 2
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))                   // 6
}
