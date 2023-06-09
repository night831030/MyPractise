package main

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	testclass := map[string]struct {
		Input  int
		Answer bool
	}{
		"class1": {
			Input:  0,
			Answer: true,
		},
		"class2": {
			Input:  12321,
			Answer: true,
		},
		"class3": {
			Input:  123321,
			Answer: true,
		},
		"class4": {
			Input:  10,
			Answer: false,
		},
	}
	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := IsPalindrome(tc.Input)
			if tc.Answer != output {
				t.Errorf("func is fail, expect %t but get %t", tc.Answer, output)
			}
		})
	}
}
