package atoi

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAtoi(t *testing.T) {
	var tests = []struct {
		inp string
		out string
	}{
		{"123", "123\n"},
		{"12ABC", "not convertable chars\n"},
		{"12123123123", "overflow\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%s, %s", tt.inp, tt.out)
		t.Run(testname, func(t *testing.T) {
			atoi(out, tt.inp)
			if out.String() != tt.out {
				t.Errorf("test for the function atoi has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.out)
			}
		})
	}
}
