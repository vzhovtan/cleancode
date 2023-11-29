package longestsubstring

import (
	"fmt"
	"testing"
)

func TestLongest(t *testing.T) {
	var tests = []struct {
		inp string
		exp int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %d", tt.inp, tt.exp)
		t.Run(testname, func(t *testing.T) {
			got := longest(tt.inp)
			if got != tt.exp {
				t.Errorf("test for the function longest has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.exp)
			}
		})
	}
}
