package main

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	testclass := map[string]struct {
		input  []int
		output []string
	}{
		"class1": {
			input:  []int{0, 1, 2, 4, 5, 7},
			output: []string{"0->2", "4->5", "7"},
		},
		"class2": {
			input:  []int{-1, 2, 3, 4, 6, 8, 9},
			output: []string{"-1", "2->4", "6", "8->9"},
		},
		"class3": {
			input:  []int{},
			output: nil,
		},
	}

	for c, tc := range testclass {
		t.Run(c, func(t *testing.T) {
			out := summaryRanges(tc.input)
			if !reflect.DeepEqual(out, tc.output) {
				t.Errorf("This func is fail, expect %s but get %s", tc.output, out)
			}
		})
	}
}
