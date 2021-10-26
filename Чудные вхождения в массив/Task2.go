package main

import (
	"fmt"
)

func main() {
	var arr = []int{2, 3, 1, 5}
	
	fmt.Println(Solution(arr))
	
}

func Solution(A []int) int {
	var number int
	
	size := len(A) + 1
	counter := 0
	
	
	for i:=1; i <= size; i++ {
		counter = 0
		for _, j := range A {
			
			if i == j {
				counter++
			}
		}
		if counter == 0 {
			number = i
		}
	}
	return number
}
