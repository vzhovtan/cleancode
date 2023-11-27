package string_in_string

import (
	"bytes"
	"fmt"
	"testing"
)

func TestStrInStr(t *testing.T) {
	var tests = []struct {
		txt string
		pat string
		exp string
	}{
		{"abracadabra", "ab", "[pattern starting index is 0 pattern starting index is 7]\n"},
		{"abracadabra", "cde", "pattern not found in string\n"},
		{"", "ab", "empty string provided\n"},
		{"abracadabra", "", "empty pattern provided\n"},
		{"", "", "empty string provided\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%s, %s,", tt.txt, tt.pat)
		t.Run(testname, func(t *testing.T) {
			strInStr(out, tt.txt, tt.pat)
			if out.String() != tt.exp {
				t.Errorf("test for the function strInStr has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.exp)
			}
		})
	}
}
