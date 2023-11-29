package longest2chars

import (
	"fmt"
	"testing"
)

func TestLongest2Chars(t *testing.T) {
	var tests = []struct {
		inp string
		exp int
	}{
		{"abaac", 4},
		{"eceba", 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %d", tt.inp, tt.exp)
		t.Run(testname, func(t *testing.T) {
			got := long2c(tt.inp)
			if got != tt.exp {
				t.Errorf("test for the function long2c has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.exp)
			}
		})
	}
}
