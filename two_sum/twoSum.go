package main

import (
	"fmt"
)

// return indexes of two slice elements which sum equal to given key
func twoSum(alist []int, akey int) {
	amap := map[int]int{}
	for i, v := range alist {
		if j, ok := amap[akey-v]; ok {
			fmt.Println(j, i)
		}
		amap[v] = i
	}
}
func main() {
	akey := 7
	alist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	twoSum(alist, akey)
}
