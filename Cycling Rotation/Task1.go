package main

import (
	"fmt"
)

func main() {
	var A = []int{3, 8, 9, 7, 6}
	fmt.Println(Solution(A[:], 3))
}

func Solution(A []int, K int) []int {
	var temp int
	
	for i:=0; i < K; i++ {
		temp = A[len(A)-1]
		
		for j:=len(A)-1; j>0; j-- {
			A[j] = A[j-1]
		}
		A[0] = temp
	}
	return A
}
