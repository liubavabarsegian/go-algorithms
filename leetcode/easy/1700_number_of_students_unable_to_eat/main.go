package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	loopCount := 0
	for len(students) > 0 && len(sandwiches) > 0 && loopCount <= len(students) {
		// если студент предпочитает бутерброд, то кушает его и уходит из очереди
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			loopCount = 0
		} else {
			student := students[0]
			students = students[1:]
			students = append(students, student)
			loopCount++
		}
	}
	return len(students)
}

// [1, 1, 1, 0, 0, 1] [1, 0, 0, 0, 1, 1]
// [1, 1, 0, 0, 1] [0, 0, 0, 1, 1] loopCount = 0
// [1, 0, 0, 1, 1] [0, 0, 0, 1, 1] loopCount = 1
// [0, 0, 1, 1, 1] [0, 0, 0, 1, 1] loopCount = 2
// [0, 1, 1, 1] [0, 0, 1, 1] loopCount = 0
// [1, 1, 1] [0, 1, 1] loopCount = 0
// [1, 1, 1] [0, 1, 1] loopCount = 1
// [1, 1, 1] [0, 1, 1] loopCount = 2
// [1, 1, 1] [0, 1, 1] loopCount = 3
func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))             //0
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1})) //3
}
