package two_sum

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		list []int
		sum  int
		exp  string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, "[elements are 2 3 elements are 1 4 elements are 0 5]\n"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 29, "elements not found\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%v, %d", tt.list, tt.sum)
		t.Run(testname, func(t *testing.T) {
			twoSum(out, tt.list, tt.sum)
			if out.String() != tt.exp {
				t.Errorf("test for the function twoSum has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.exp)
			}
		})
	}
}

func TestTwoSumSorted(t *testing.T) {
	var tests = []struct {
		list []int
		sum  int
		exp  string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, "[elements are 5 0 elements are 4 1 elements are 3 2]\n"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 29, "elements not found\n"},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		testname := fmt.Sprintf("%v, %d", tt.list, tt.sum)
		t.Run(testname, func(t *testing.T) {
			twoSumSorted(out, tt.list, tt.sum)
			if out.String() != tt.exp {
				t.Errorf("test for the function twoSum has failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), tt.exp)
			}
		})
	}
}
