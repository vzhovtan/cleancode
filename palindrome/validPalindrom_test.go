package palindrome

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPalindrom(t *testing.T) {
	var tests = []struct {
		in  string
		exp string
	}{
		{"A man, a plan, a canal: Panama", "Valid palindrom\n"},
		{"", "Valid palindrom\n"},
		{"Example one", "Not a valid palindrom\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%s, %s,", tt.in, tt.exp)
		t.Run(testname, func(t *testing.T) {
			palindrome(out, tt.in)
			if out.String() != tt.exp {
				t.Errorf("test for the function palindrome has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.exp)
			}
		})
	}
}
