package problem_002

import "math"

func nthFib(n float64) float64 {
	return math.Round((math.Pow(math.Phi, n) - math.Pow((1-math.Phi), n)) / math.Sqrt(5))
}

func solveProblem(limit float64) float64 {
	i := 3
	sum := 0.0
	for {
		val := nthFib(float64(i))
		if val > limit {
			break
		}
		sum += val
		i += 3
	}

	return sum
}
