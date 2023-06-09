package main

import "testing"

func TestLemonadeChange(t *testing.T) {
	testclass := map[string]struct {
		input  []int
		output bool
	}{
		"class1": {
			input:  []int{5, 5, 10, 10, 20},
			output: false,
		},
		"class2": {
			input:  []int{5, 5, 5, 10, 10, 5, 20, 20},
			output: true,
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := lemonadeChange(tc.input)
			if output != tc.output {
				t.Errorf("func is fail, expect %v but get %v", tc.output, output)
			}
		})
	}

}
