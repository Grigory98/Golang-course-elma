package main

import "fmt"

func main() {
	var arr = []int{4, 1, 3, 2}
	fmt.Println(Solution(arr))
}

func Solution(A []int) int {
	var temp int
	temp = 0

	for i := 1; i <= len(A); i++ {

		if !findElem(A, i) {
			temp = temp + 1
		}
	}

	if temp > 0 {
		return 0
	}
	return 1
}

/**
* Проверяет вхождения в массив
**/
func findElem(A []int, x int) bool {
	for _, n := range A {
		if x == n {
			return true
		}
	}
	return false
}
