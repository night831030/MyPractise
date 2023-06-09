package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testclass := map[string]struct {
		input  string
		output int
	}{
		"class1": {
			input:  "Hello World",
			output: 5,
		},
		"class2": {
			input:  "   fly me   to   the moon  ",
			output: 4,
		},
		"class3": {
			input:  "luffy is still joyboy",
			output: 6,
		},
	}

	for name, tclass := range testclass {
		output := lengthOfLastWord(tclass.input)
		t.Run(name, func(t *testing.T) {
			if output != tclass.output {
				t.Errorf("expect %d but get %d, func is fail", tclass.output, output)
			}
		})
	}
}
