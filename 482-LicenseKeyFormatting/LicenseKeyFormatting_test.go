package main

import "testing"

func TestLicenseKeyFormatting(t *testing.T) {
	testclass := map[string]struct {
		input1 string
		input2 int
		output string
	}{
		"class1": {
			input1: "5F3Z-2e-9-w",
			input2: 4,
			output: "5F3Z-2E9W",
		},
		"class2": {
			input1: "2-5g-3-J",
			input2: 2,
			output: "2-5G-3J",
		},
		"class3": {
			input1: "2-4A0r7-4k",
			input2: 4,
			output: "24A0-R74K",
		},
	}

	for name, tclass := range testclass {
		t.Run(name, func(t *testing.T) {
			output := licenseKeyFormatting(tclass.input1, tclass.input2)
			if output != tclass.output {
				t.Errorf("expect %s get %s,func is fail", tclass.output, output)
			}
		})
	}
}
