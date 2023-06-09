package main

import "testing"

func TestLargestTriangleArea(t *testing.T) {
	testclass := map[string]struct {
		input  [][]int
		output float64
	}{
		"class1": {
			input:  [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			output: 2.00000,
		},
		"class2": {
			input:  [][]int{{1, 0}, {0, 0}, {0, 1}},
			output: 0.50000,
		},
	}

	for n, tc := range testclass {
		out := largestTriangleArea(tc.input)
		t.Run(n, func(t *testing.T) {
			if out != tc.output {
				t.Errorf("func is fail,expect %f but get %f", tc.output, out)
			}
		})
	}
}
