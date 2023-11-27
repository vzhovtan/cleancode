package palindrom

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// finding if the given string is valid palindrom with two pointers
func palindrome(w io.Writer, astring string) {
	var b strings.Builder
	for _, r := range astring {
		if unicode.IsLetter(r) {
			fmt.Fprintf(&b, strings.ToLower(string(r)))
		}
	}

	newString := b.String()

	for i, j := 0, len(newString)-1; i < j; {
		if newString[i] != newString[j] {
			fmt.Fprintf(w, "Not a valid palindrom\n")
			return
		}
		i++
		j--
	}
	fmt.Fprintf(w, "Valid palindrom\n")
}

func main() {
	astring := "A man, a plan, a canal: Panama"
	palindrome(os.Stdout, astring)
}
