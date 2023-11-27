package two_sum_sorted

import (
	"fmt"
	"io"
	"os"
)

// return indexes of two sorted slice elements which sum equal to given key
// two pointers approach
func twoSum(w io.Writer, alist []int, akey int) {
	res := []string{}
	for i, j := 0, len(alist)-1; i < j; {
		if alist[i]+alist[j] > akey {
			j--
		} else if alist[i]+alist[j] < akey {
			i++
		} else {
			el := fmt.Sprintf("elements are %d %d", j, i)
			res = append(res, el)
			i++ // increasing firs pointer to be sure there is no endless loop, can decrease second pointer as well
		}
	}
	if len(res) != 0 {
		fmt.Fprintf(w, "%v\n", res)
	} else {
		fmt.Fprintf(w, "elements not found\n")
	}
}

func main() {
	alist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	akey := 7
	twoSum(os.Stdout, alist, akey)
}
