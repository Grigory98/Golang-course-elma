package main

import "fmt"

func main() {
	var arr = []int{6, 3, 5, 4}
	fmt.Println(Solution(arr))
}

func Solution(A []int) int {
	var temp int
	temp = 0
	min := findMin(A)
	for i := 1; i <= len(A); i++ {

		if !findElem(A, min) {
			temp = temp + 1
		}
		min = min + 1
	}

	if temp > 0 {
		return 0
	}
	return 1
}

//Проверяет вхождения в массив
func findElem(A []int, x int) bool {
	for _, n := range A {
		if x == n {
			return true
		}
	}
	return false
}

//Находит минимальное значение в массиве
func findMin(values []int) int {
	min := values[0]

	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
