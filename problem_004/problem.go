package problem_004

func reverseNumber(n int) int {
	rev := 0
	for n > 0 {
		r := n % 10
		rev *= 10
		rev += r
		n /= 10
	}

	return rev
}

func isPalindrome(n int) bool {
	r := reverseNumber(n)

	return r == n
}

func solveProblem(lowest int, highest int) int {
	max := 0
	for i := lowest; i <= highest; i++ {
		for y := lowest; y <= highest; y++ {
			n := i * y
			if isPalindrome(n) && n > max {
				max = n
			}
		}
	}

	return max
}
