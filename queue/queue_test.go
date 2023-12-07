package queue

import (
	"fmt"
	"testing"
)

func TestEnqueueDequeue(t *testing.T) {
	var tests = []struct {
		ip []int
		ot []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.ip, tt.ot)
		t.Run(testname, func(t *testing.T) {
			q := queue{
				head: nil,
				len:  0,
			}
			got := []int{}
			for _, v := range tt.ip {
				q.enqueue(v)
			}

			for _, _ = range got {
				got = append(got, q.dequeue())
			}

			pass := true
			for i := 0; i < len(got); i++ {
				if tt.ot[i] != got[i] {
					pass = false
				}
			}

			if !pass {
				t.Errorf("test for the function enqueue and dequeue has failed - results not match\nGot:\n%v\nExpected:\n%v", got, tt.ot)
			}
		})
	}
}
