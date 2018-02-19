package problem_002

import "math"

func nthFib(input float64) float64 {
	return math.Round((math.Pow(math.Phi, input) - math.Pow((1-math.Phi), input)) / math.Sqrt(5))
}

func solveProblem(cap float64) float64 {
	counter := 3
	sum := 0.0
	for {
		val := nthFib(float64(counter))
		if val > cap {
			break
		}
		sum += val
		counter += 3
	}

	return sum
}
