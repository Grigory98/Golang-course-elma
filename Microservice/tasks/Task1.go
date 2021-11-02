//Циклическая ротация
package tasks

//Cycling rotation
func SolutionTask1(A []int, K int) []int {
	var temp int

	for i := 0; i < K; i++ {
		temp = A[len(A)-1]

		for j := len(A) - 1; j > 0; j-- {
			A[j] = A[j-1]
		}
		A[0] = temp
	}
	return A
}
