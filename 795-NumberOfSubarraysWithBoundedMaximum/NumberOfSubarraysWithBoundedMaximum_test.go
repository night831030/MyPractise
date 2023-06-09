package main

import "testing"

func TestNumSubarrayBoundedMax(t *testing.T) {
	testclass := map[string]struct {
		input1 []int
		input2 int
		input3 int
		output int
	}{
		"class1": {
			input1: []int{2, 1, 4, 3},
			input2: 2,
			input3: 3,
			output: 3,
		},
		"class2": {
			input1: []int{2, 9, 2, 5, 6},
			input2: 2,
			input3: 8,
			output: 7,
		},
		"class3": {
			input1: []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52},
			input2: 32,
			input3: 69,
			output: 22,
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := numSubarrayBoundedMax(tc.input1, tc.input2, tc.input3)
			if tc.output != output {
				t.Errorf("function is fail, expect %d but get %d", tc.output, output)
			}
		})
	}

}
