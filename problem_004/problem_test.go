package problem_004

import "testing"

type TestCase struct {
	lowest   int
	highest  int
	expected int
}

var tests = []TestCase{
	{100, 999, 906609},
}

func TestSolveProblem(t *testing.T) {
	for _, test := range tests {
		result := solveProblem(test.lowest, test.highest)
		if result != test.expected {
			t.Errorf("Expected [%v], got [%v] for input [%v]-[%v]", test.expected, result, test.lowest, test.highest)
		}
	}
}
