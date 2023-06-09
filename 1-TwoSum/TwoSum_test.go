package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testclass := map[string]struct {
		input1 []int
		input2 int
		answer []int
	}{
		"class1": {
			input1: []int{9, 5, 3, 4},
			input2: 12,
			answer: []int{0, 2},
		},
		"class2": {
			input1: []int{5, 9, 11, 78},
			input2: 87,
			answer: []int{1, 3},
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := TwoSum(tc.input1, tc.input2)
			if !reflect.DeepEqual(tc.answer, output) {
				t.Errorf("func is fail, expect %d but get %d", tc.answer, output)
			}
		})

	}
}
