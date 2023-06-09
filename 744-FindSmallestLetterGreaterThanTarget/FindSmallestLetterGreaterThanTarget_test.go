package main

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	testclass := map[string]struct {
		input1 []byte
		input2 byte
		output byte
	}{
		"class1": {
			input1: []byte{'c', 'f', 'j'},
			input2: 'a',
			output: 'c',
		},
		"class2": {
			input1: []byte{'c', 'f', 'j'},
			input2: 'c',
			output: 'f',
		},
		"class3": {
			input1: []byte{'x', 'x', 'y', 'y'},
			input2: 'z',
			output: 'x',
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := nextGreatestLetter(tc.input1, tc.input2)
			if output != tc.output {
				t.Errorf("func is fail, expect %d but get %d", tc.output, output)
			}
		})
	}
}
