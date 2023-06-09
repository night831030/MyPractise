package main

import "testing"

func TestFindNearesPoint(t *testing.T) {
	testclass := map[string]struct {
		input1 int
		input2 int
		input3 [][]int
		answer int
	}{
		"class1": {
			input1: 5,
			input2: 10,
			input3: [][]int{{3, 5}, {5, 7}, {1000, 10}, {5, -300}, {5, 10}},
			answer: 4,
		},
		"class2": {
			input1: 3,
			input2: 4,
			input3: [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}},
			answer: 2,
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := NearestValidPoint(tc.input1, tc.input2, tc.input3)
			if tc.answer != output {
				t.Errorf("func is fail, expect %d but get %d", tc.answer, output)
			}
		})
	}
}
