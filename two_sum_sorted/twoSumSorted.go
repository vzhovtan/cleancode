package main

import "fmt"

// return indexes of two sorted slice elements which sum equal to given key
// two pointers approach
func twoSum(alist []int, akey int) {
	for i, j := 0, len(alist)-1; i < j; {
		if alist[i]+alist[j] > akey {
			j--
		} else if alist[i]+alist[j] < akey {
			i++
		} else {
			fmt.Println(i, j)
			i++ // increasing firs pointer to be sure there is no endless loop, can decrease second pointer as well
		}
	}
}

func main() {
	alist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	akey := 7
	twoSum(alist, akey)
}
