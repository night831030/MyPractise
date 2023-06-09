package main

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	testclass := map[string]struct {
		Input1 []byte
		Input2 byte
		output byte
	}{
		"class1": {
			Input1: []byte{'c', 'f', 'j'},
			Input2: 'a',
			output: 'c',
		},
		"class2": {
			Input1: []byte{'c', 'f', 'j'},
			Input2: 'c',
			output: 'f',
		},
		"class3": {
			Input1: []byte{'x', 'x', 'y', 'y'},
			Input2: 'z',
			output: 'x',
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := nextGreatestLetter(tc.Input1, tc.Input2)
			if output != tc.output {
				t.Errorf("func is fail, expect %d but get %d", tc.output, output)
			}
		})
	}
}
