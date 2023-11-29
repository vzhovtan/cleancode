package missingranges

import (
	"fmt"
	"testing"
)

func TestPalindrom(t *testing.T) {
	var tests = []struct {
		in  []int
		exp string
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, "{}"},
		{[]int{}, "{0->9}"},
		{[]int{0, 1, 3, 5, 66, 89}, "{2, 4, 6 -> 65, 67 -> 88}"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %s,", tt.in, tt.exp)
		t.Run(testname, func(t *testing.T) {
			got := findMissing(tt.in)
			if got != tt.exp {
				t.Errorf("test for the function findMissing has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.exp)
			}
		})
	}
}
