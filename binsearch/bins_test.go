package binsearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		ip []int
		lo int
		n  int
		ot bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 4, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 6, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 11, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.ip, tt.ot)
		t.Run(testname, func(t *testing.T) {
			hi := len(tt.ip)
			got := search(tt.ip, tt.lo, hi, tt.n)
			if got != tt.ot {
				t.Errorf("test for the function search has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.ot)
			}
		})
	}
}
