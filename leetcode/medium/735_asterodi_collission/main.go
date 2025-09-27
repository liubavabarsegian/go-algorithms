package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, v := range asteroids {
		// если астероид летит вправо, то записываем его
		if v > 0 || len(stack) == 0 {
			stack = append(stack, v)
			continue

		}

		// дальше астероиды летят только влево
		// если прошлый астероид летит в том же направлении, то записываем текущий
		if stack[len(stack)-1] < 0 {
			stack = append(stack, v)
			continue
		}

		// если астероид летит влево, он может столкнуться с другими
		// пока на встречу текущему астероиду, летят другие
		// он готов с ними сталкиваться (если они меньше его)
		for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] <= -v {
			// если размеры астероидов равны, ликвидируем оба
			if stack[len(stack)-1]*-1 == v {
				stack = stack[0 : len(stack)-1]
				break
				// если слева астероид меньше текущего, то мы его сбиваем
			} else if stack[len(stack)-1] < -v {
				stack = stack[0 : len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, v)
			}
		}

	}
	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))     // [5, 10]
	fmt.Println(asteroidCollision([]int{8, -8}))         // []
	fmt.Println(asteroidCollision([]int{10, 2, -5}))     // [10]
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))  // [-2, -1, 1, 2]
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2})) // [-2, -2, -2]
	fmt.Println(asteroidCollision([]int{1, -2, -2, -2})) // [-2, -2, -2]
}
