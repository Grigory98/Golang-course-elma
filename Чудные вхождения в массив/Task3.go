package main

import "fmt"

func main() {
	var arr = []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println(Solution(arr))
}

func Solution(A []int) int {
	var temp int

	for i := 0; i < len(A); i++ {

		if i < len(A)-2 {

			if A[i] != A[i+2] {
				temp = A[i+2]
			}

		}

	}
	return temp
}
