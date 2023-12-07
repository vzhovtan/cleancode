//Package validnumber is verifying if the provided string is valid decimal number
package validnumber

import (
	"fmt"
	"io"
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

}
