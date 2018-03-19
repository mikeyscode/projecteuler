package problem_005

func solveProblem(low, high int) int {
	k := 1
	numbers := makeRange(low, high)
	for {
		if ldn(k, numbers) == true {
			break
		}

		k++
	}

	return k
}

func ldn(number int, numbers []int) bool {
	for _, n := range numbers {
		if number%n != 0 {
			return false
		}
	}

	return true
}

func makeRange(min, max int) []int {
	numbers := make([]int, max-min+1)
	for i := range numbers {
		numbers[i] = min + i
	}

	return numbers
}
