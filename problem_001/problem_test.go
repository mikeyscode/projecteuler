package problem_001

import "testing"

type TestCase struct {
	x        int
	y        int
	max      int
	expected int
}

var tests = []TestCase{
	{3, 5, 10, 23},
	{3, 5, 100, 2318},
	{3, 5, 1000, 233168},
	{3, 5, 10000, 23331668},
}

func TestSolveProblem(t *testing.T) {
	// 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	for _, test := range tests {
		result := solveProblem(test.x, test.y, test.max)
		if result != test.expected {
			t.Errorf("Expected [%d], got [%d] for inputs: max[%d], x[%d] and y[%d]", test.expected, result, test.max, test.x, test.y)
		}
	}
}
