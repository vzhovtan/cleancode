package main

import (
	"fmt"
	"strings"
	"unicode"
)

// finding if the given string is valid palindrom with two pointers
func plaindrome(astring string) {
	var b strings.Builder
	for _, r := range astring {
		if unicode.IsLetter(r) {
			fmt.Fprintf(&b, strings.ToLower(string(r)))
		}
	}

	newString := b.String()

	for i, j := 0, len(newString)-1; i < j; {
		if newString[i] != newString[j] {
			fmt.Println("Not a vlaid palindrom")
		}
		i++
		j--
	}
	fmt.Println("Valid palindrom")
}

func main() {
	astring := "A man, a plan, a canal: Panama"
	plaindrome(astring)
}
