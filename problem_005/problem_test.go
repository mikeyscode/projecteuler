package problem_005

import "testing"

type TestCase struct {
	low      int
	high     int
	expected int
}

var tests = []TestCase{
	{1, 10, 2520},
	{1, 20, 232792560},
}

func TestSolveProblem(t *testing.T) {
	for _, test := range tests {
		result := solveProblem(test.low, test.high)
		if result != test.expected {
			t.Errorf("Expected [%v], got [%v] for input [%v] and [%v]", test.expected, result, test.low, test.high)
		}
	}
}
