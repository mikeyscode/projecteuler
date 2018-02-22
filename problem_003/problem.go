package problem_003

import "math"

func isPrime(n int64) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}

	var i int64 = 5
	var w int64 = 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}

	return true
}

func solveProblem(n float64) float64 {
	p := math.Round(math.Sqrt(n) + 0.5)
	for p > 0 {
		if isPrime(int64(p)) && int64(n)%int64(p) == 0 {
			return p
		}
		p--
	}
	return 0
}
