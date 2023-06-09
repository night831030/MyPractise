package main

import "testing"

func TestNumSubarrayBoundedMax(t *testing.T) {
	testclass := map[string]struct {
		Input1 []int
		Input2 int
		Input3 int
		output int
	}{
		"class1": {
			Input1: []int{2, 1, 4, 3},
			Input2: 2,
			Input3: 3,
			output: 3,
		},
		"class2": {
			Input1: []int{2, 9, 2, 5, 6},
			Input2: 2,
			Input3: 8,
			output: 7,
		},
		"class3": {
			Input1: []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52},
			Input2: 32,
			Input3: 69,
			output: 22,
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := numSubarrayBoundedMax(tc.Input1, tc.Input2, tc.Input3)
			if tc.output != output {
				t.Errorf("function is fail, expect %d but get %d", tc.output, output)
			}
		})
	}

}
