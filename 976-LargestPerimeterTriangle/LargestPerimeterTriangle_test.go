package main

import "testing"

func TestLargestPerimeter(t *testing.T) {
	testclass := map[string]struct {
		input  []int
		output int
	}{
		"class1": {
			input:  []int{2, 6, 3, 2, 3},
			output: 8,
		},
		"class2": {
			input:  []int{9, 4, 1, 2},
			output: 0,
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := LargestPerimeter(tc.input)
			if output != tc.output {
				t.Errorf("func is fail, expect %d but get %d", tc.output, output)
			}
		})
	}

}
