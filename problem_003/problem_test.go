package problem_003

import "testing"

type TestCase struct {
	input    float64
	expected float64
}

var tests = []TestCase{
	{600851475143, 6857},
}

func TestSolveProblem(t *testing.T) {
	for _, test := range tests {
		result := solveProblem(test.input)
		if result != test.expected {
			t.Errorf("Expected [%v], got [%v] for input [%v]", test.expected, result, test.input)
		}
	}
}
