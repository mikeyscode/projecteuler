package problem_006

import "testing"

type TestCase struct {
	cap      int
	expected int
}

var tests = []TestCase{
	{100, 25164150},
}

func TestSolveProblem(t *testing.T) {
	for _, test := range tests {
		result := solveProblem(test.cap)
		if result != test.expected {
			t.Errorf("Expected [%v], got [%v] for input [%v]", test.expected, result, test.cap)
		}
	}
}
