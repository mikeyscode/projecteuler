package problem_001

func solveProblem(x, y, max int) int {
	lcm := calculateLowestCommonMultiple(x, y)
	sumX := calculateSumUptoNumber(x, max)
	sumY := calculateSumUptoNumber(y, max)
	sumLCM := calculateSumUptoNumber(lcm, max)

	return sumX + sumY - sumLCM
}

// calculateLeastCommonMultiple works out the lowest common multiple for two numbers
func calculateLowestCommonMultiple(x, y int) int {
	multiple := x
	if x < y {
		multiple = y
	}

	for {
		if multiple%x == 0 && multiple%y == 0 {
			return multiple
		}

		multiple++
	}
}

// calculateSumUptoNumber works out the sum of numbers up to maximum amount divided by n
// the sum of all numbers less than n that are divisible by d is calcualted by:
// n * (max-1 / n) * ((max-1 / n)+1) / 2
func calculateSumUptoNumber(n, max int) int {
	return n * ((max - 1) / n) * (((max - 1) / n) + 1) / 2
}
