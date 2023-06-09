package main

import (
	"reflect"
	"testing"
)

func TestBinaryWatch(t *testing.T) {
	testclass := map[string]struct {
		input  int
		answer []string
	}{
		"class1": {
			input:  1,
			answer: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		"class2": {
			input:  8,
			answer: []string{"7:31", "7:47", "7:55", "7:59", "11:31", "11:47", "11:55", "11:59"},
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := readBinaryWatch(tc.input)
			if !reflect.DeepEqual(output, tc.answer) {
				t.Errorf("func is fail, expect %s but get %s", tc.answer, output)
			}
		})
	}
}
