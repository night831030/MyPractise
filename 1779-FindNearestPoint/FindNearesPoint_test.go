package main

import "testing"

func TestFindNearesPoint(t *testing.T) {
	testclass := map[string]struct {
		Input1 int
		Input2 int
		Input3 [][]int
		Answer int
	}{
		"class1": {
			Input1: 5,
			Input2: 10,
			Input3: [][]int{{3, 5}, {5, 7}, {1000, 10}, {5, -300}, {5, 10}},
			Answer: 4,
		},
		"class2": {
			Input1: 3,
			Input2: 4,
			Input3: [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}},
			Answer: 2,
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := NearestValidPoint(tc.Input1, tc.Input2, tc.Input3)
			if tc.Answer != output {
				t.Errorf("func is fail, expect %d but get %d", tc.Answer, output)
			}
		})
	}
}
