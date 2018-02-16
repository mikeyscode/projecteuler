package problem_001

// solveProblem will list all the natural numbers below a given input that are multiples of 3 or 5 and then
// sum of these multiples is 23.
func solveProblem(input int) int {
	var sum int
	for i := 1; i < input; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
