package main

import (
	"math/bits"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	testclass := map[string]struct {
		Input  uint32
		Answer int
	}{
		"class1": {
			Input:  7,
			Answer: bits.OnesCount32(7),
		},
		"class2": {
			Input:  0,
			Answer: bits.OnesCount32(0),
		},
		"class3": {
			Input:  32,
			Answer: bits.OnesCount32(32),
		},
	}

	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := HammingWeight(tc.Input)
			if output != tc.Answer {
				t.Errorf("func is fail, expect %d but get %d", tc.Answer, output)
			}
		})
	}

}
