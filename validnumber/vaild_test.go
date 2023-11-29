package validnumber

import (
	"bytes"
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	var tests = []struct {
		inp string
		out string
	}{
		{"123", "valid number\n"},
		{"12ABC", "not a valid number\n"},
		{"121.23123123", "valid number\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%s, %s", tt.inp, tt.out)
		t.Run(testname, func(t *testing.T) {
			valid(out, tt.inp)
			if out.String() != tt.out {
				t.Errorf("test for the function valid has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.out)
			}
		})
	}
}
