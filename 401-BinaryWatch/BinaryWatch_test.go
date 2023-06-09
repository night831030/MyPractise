package main

import (
	"reflect"
	"testing"
)

func TestBinaryWatch(t *testing.T) {
	testclass := map[string]struct {
		Input  int
		Answer []string
	}{
		"class1": {
			Input:  1,
			Answer: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		"class2": {
			Input:  8,
			Answer: []string{"7:31", "7:47", "7:55", "7:59", "11:31", "11:47", "11:55", "11:59"},
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := readBinaryWatch(tc.Input)
			if !reflect.DeepEqual(output, tc.Answer) {
				t.Errorf("func is fail, expect %s but get %s", tc.Answer, output)
			}
		})
	}
}
