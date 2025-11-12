package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	if n <= 1 {
		return -1
	}

	left := 0
	right := n - 1
	for left <= right {
		medium := (left + right) / 2
		res := guess(medium)
		switch res {
		case 1:
			right = medium - 1
		case -1:
			left = medium + 1
		case 0:
			return medium
		}
	}
	return -1
}

func main() {

}
