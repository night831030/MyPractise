package main

import (
	"math/bits"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	testclass := map[string]struct {
		input  uint32
		answer int
	}{
		"class1": {
			input:  7,
			answer: bits.OnesCount32(7),
		},
		"class2": {
			input:  0,
			answer: bits.OnesCount32(0),
		},
		"class3": {
			input:  32,
			answer: bits.OnesCount32(32),
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := HammingWeight(tc.input)
			if output != tc.answer {
				t.Errorf("func is fail, expect %d but get %d", tc.answer, output)
			}
		})
	}

}
