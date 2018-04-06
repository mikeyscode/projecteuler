package problem_006

func solveProblem(cap int) int {
	squaredSum := 0
	sumSquared := 0
	for i := 1; i <= cap; i++ {
		squaredSum += i * i
		sumSquared += i
	}
	sumSquared *= sumSquared

	return sumSquared - squaredSum
}
