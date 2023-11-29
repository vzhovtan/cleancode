//Package atoi returns the string converted to a decimal integer
package atoi

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

var maxNum = 214748364

//check if the chars is digit and if there will be on overlofw for int32
//skipping space trimming and number sign verification for simplicity
func atoi(w io.Writer, inp string) {
	num := 0
	for _, v := range inp {
		if unicode.IsDigit(v) {
			dig, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatal(err)
			}
			if num > maxNum || (num == maxNum && dig > 8) {
				fmt.Fprintf(w, "overflow\n")
				return
			}
			num = num*10 + dig
		} else {
			fmt.Fprintf(w, "not convertable chars\n")
			return
		}
	}
	fmt.Fprintf(w, "%d\n", num)
}

func main() {
	atoi(os.Stdout, "1213")
	atoi(os.Stdout, "12130000000")
	atoi(os.Stdout, "12abc")
}
