package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testclass := map[string]struct {
		Input1 []int
		Input2 int
		Answer []int
	}{
		"class1": {
			Input1: []int{9, 5, 3, 4},
			Input2: 12,
			Answer: []int{0, 2},
		},
		"class2": {
			Input1: []int{5, 9, 11, 78},
			Input2: 87,
			Answer: []int{1, 3},
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := TwoSum(tc.Input1, tc.Input2)
			if !reflect.DeepEqual(tc.Answer, output) {
				t.Errorf("func is fail, expect %d but get %d", tc.Answer, output)
			}
		})

	}
}
