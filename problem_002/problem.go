package problem_002

func evenFibSum(cap int) int {
	sum, x, y := 0, 0, 1

	for y < cap {
		y, x = x+y, y

		if y%2 == 0 {
			sum += y
		}
	}

	return sum
}

func solveProblem(cap int) int {
	return evenFibSum(cap)
}
