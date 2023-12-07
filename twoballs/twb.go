// Package twoballs implements solution for two crystal balls problem
package twoballs

import (
	"math"
)

func twb(arr []bool) int {
	jmp := int(math.Sqrt(float64(len(arr))))
	a := 0

	if len(arr) == 0 {
		return -1
	}

	if arr[0] {
		return 0
	}

	for i := 0; i < len(arr); i += jmp {
		if arr[i] {
			a = i
			break
		}
	}

	if a == 0 {
		for k := len(arr) - jmp; k < len(arr); k++ {
			if arr[k] {
				return k
			}
		}
	} else {
		for j := a - jmp; j < a+jmp; j++ {
			if arr[j] {
				return j
			}
		}
	}
	return -1
}

func main() {

}
