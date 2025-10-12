package main

import "fmt"

/*
	_1_
		_1_
			_2_
				_3_
					_7_
*/

/*
	_1_
		_1_
			_2_
				_4_
					_7_
						_13_
*/

func countPaths(stairs uint) uint {
	stairsPaths := make([]uint, stairs+1)

	stairsPaths[0] = 1
	stairsPaths[1] = 1
	if stairs > 1 {
		stairsPaths[2] = 2
		for i := uint(3); i < stairs+1; i++ {
			stairsPaths[i] = stairsPaths[i-1] + stairsPaths[i-2] + stairsPaths[i-3]
		}
	}

	return stairsPaths[len(stairsPaths)-1]
}

func main() {
	var stairs uint

	fmt.Scanf("%d\n", &stairs)
	fmt.Println(countPaths(stairs))
}
