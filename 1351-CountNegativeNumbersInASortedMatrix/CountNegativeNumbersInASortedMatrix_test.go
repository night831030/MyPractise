package main

import "testing"

func TestCountNegatives(t *testing.T) {
	testclass := map[string]struct {
		Input  [][]int
		output int
	}{
		"class1": {
			Input:  [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			output: 8,
		},
		"class2": {
			Input:  [][]int{{3, 2}, {1, 0}},
			output: 0,
		},
		"class3": {
			Input:  [][]int{{5, 1, 0}, {-5, -5, -5}},
			output: 3,
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := countNegatives(tc.Input)
			if output != tc.output {
				t.Errorf("func is fail, expect %d but get %d", tc.output, output)
			}
		})
	}
}
