package main

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	testclass := map[string]struct {
		input  int
		answer bool
	}{
		"class1": {
			input:  0,
			answer: true,
		},
		"class2": {
			input:  12321,
			answer: true,
		},
		"class3": {
			input:  123321,
			answer: true,
		},
		"class4": {
			input:  10,
			answer: false,
		},
	}
	for name, tc := range testclass {
		t.Run(name, func(t *testing.T) {
			output := IsPalindrome(tc.input)
			if tc.answer != output {
				t.Errorf("func is fail, expect %t but get %t", tc.answer, output)
			}
		})
	}
}
