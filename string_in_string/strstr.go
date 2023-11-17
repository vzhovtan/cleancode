package main

import (
	"fmt"
)

// find pattern in text and print index for each instance. naive methods with 2 pointers
func strInStr(txt, pat string) {
	for i := 0; i < len(txt)-1; {
		for j := 0; j <= len(pat)-1; {
			if txt[i] == pat[j] {
				if j == len(pat)-1 {
					fmt.Println("found, index is ", i-j)
				}
				j++
				i++
			} else if txt[i] != pat[j] {
				break
			}
		}
		i++
	}
}

func main() {
	txt := "abracadabra"
	pat := "ab"
	strInStr(txt, pat)
}
