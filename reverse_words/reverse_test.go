package reverse_words

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReverseNaive(t *testing.T) {
	var tests = []struct {
		inp string
		out string
	}{
		{"the sky is blue", "blue is sky the\n"},
		{"blue", "blue\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%s, %s", tt.inp, tt.out)
		t.Run(testname, func(t *testing.T) {
			reverseNaive(out, tt.inp)
			if out.String() != tt.out {
				t.Errorf("test for the function reverseNaive has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.out)
			}
		})
	}
}

func TestReverseOnePass(t *testing.T) {
	var tests = []struct {
		inp string
		out string
	}{
		{"the sky is blue", "blue is sky the\n"},
		{"blue", "blue\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%s, %s", tt.inp, tt.out)
		t.Run(testname, func(t *testing.T) {
			reverseOnePass(out, tt.inp)
			if out.String() != tt.out {
				t.Errorf("test for the function reverseOnePass has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.out)
			}
		})
	}
}
