package twoballs

import (
	"fmt"
	"testing"
)

func TestTwoBalls(t *testing.T) {
	var tests = []struct {
		ip []bool
		ot int
	}{
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}, 20},
		{[]bool{}, -1},
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}, 20},
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}, -1},
		{[]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.ip, tt.ot)
		t.Run(testname, func(t *testing.T) {
			got := twb(tt.ip)
			if got != tt.ot {
				t.Errorf("test for the function twb has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.ot)
			}
		})
	}
}
