package validnumber

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

//check if each chars is digit except possible decimal point
//skipping space trimming and number sign verification for simplicity
func valid(w io.Writer, inp string) {
	for _, v := range inp {
		if !unicode.IsDigit(v) && string(v) != "." {
			fmt.Fprintf(w, "not a valid number\n")
			return
		}
	}
	fmt.Fprintf(w, "valid number\n")
}

func main() {
	valid(os.Stdout, "1213")
	valid(os.Stdout, "121.3")
	valid(os.Stdout, "12abc")
}
