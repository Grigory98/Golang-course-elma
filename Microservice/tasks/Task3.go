package tasks

func SolutionTask3(A []int) int {
	var temp int
	size := len(A)
	counter := 0

	for i := 0; i < size; i++ {
		counter = 0
		for j := 0; j < size; j++ {
			if A[i] == A[j] {
				counter = counter + 1
			}
		}

		if counter%2 != 0 {
			temp = A[i]
		}

	}
	return temp
}
