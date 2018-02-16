package problem_001

import "testing"

type TestCase struct {
	input    int
	expected int
}

var tests = []TestCase{
	{10, 23},
	{100, 2318},
	{1000, 233168},
	{10000, 23331668},
}

func TestSolveProblem(t *testing.T) {
	// 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	for _, test := range tests {
		result := solveProblem(test.input)
		if result != test.expected {
			t.Errorf("Expected [%d], got [%d] for input [%d]", test.expected, result, test.input)
		}
	}
}
