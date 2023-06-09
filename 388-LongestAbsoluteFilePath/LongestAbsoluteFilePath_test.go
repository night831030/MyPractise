package main

import (
	"testing"
)

func TestLengthLongestPath(t *testing.T) {
	testclass := map[string]struct {
		input  string
		output int
	}{
		"class1": {
			input:  "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			output: 20,
		},
		"class2": {
			input:  "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			output: 32,
		},
		"class3": {
			input:  "a",
			output: 0,
		},
		"class4": {
			input:  "file1.txt\nfile2.txt\nlongfile.txt",
			output: 12,
		},
		"class5": {
			input:  "dir\n\t        file.txt\n\tfile2.txt",
			output: 20,
		},
		"class6": {
			input:  "a\n\tb.txt\na2\n\tb2.txt",
			output: 9,
		},
		"class7": {
			input:  "a\n\taa\n\t\taaa\n\t\t\tfile1.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png",
			output: 29,
		},
	}

	for name, tc := range testclass {
		output := lengthLongestPath(tc.input)
		t.Run(name, func(t *testing.T) {
			if output != tc.output {
				t.Errorf("expect %d but get %d, func is fail", tc.output, output)
			}
		})
	}

}
