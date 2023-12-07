package bubsort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	var tests = []struct {
		ip []int
		ot []int
	}{
		{[]int{3, 2, 1, 4, 5, 6, 7, 9, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{3, 4, 5, 1, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.ip, tt.ot)
		t.Run(testname, func(t *testing.T) {
			got := sort(tt.ip)
			pass := true
			for i := 0; i < len(got); i++ {
				if tt.ot[i] != got[i] {
					pass = false
				}
			}
			if !pass {
				t.Errorf("test for the function sort has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.ot)
			}
		})
	}
}
