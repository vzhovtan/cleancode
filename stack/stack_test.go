package stack

import (
	"fmt"
	"testing"
)

func TestPushPop(t *testing.T) {
	var tests = []struct {
		ip []int
		ot []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.ip, tt.ot)
		t.Run(testname, func(t *testing.T) {
			st := Stack{}
			got := []int{}
			for _, v := range tt.ip {
				st.push(v)
			}
			for i := 0; i < len(tt.ip); i++ {
				got = append(got, st.pop())
			}
			pass := true
			for i := 0; i < len(got); i++ {
				if tt.ot[i] != got[i] {
					pass = false
				}
			}
			if !pass {
				t.Errorf("test for the function poush and pop has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.ot)
			}
		})
	}
}
