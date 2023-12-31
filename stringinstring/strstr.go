//Package stringinstring finds pattern in text and print index for each instance using naive methods with 2 pointers
package stringinstring

import (
	"fmt"
	"io"
)

func strInStr(w io.Writer, txt, pat string) {
	var ind string
	var res []string

	if len(txt) == 0 {
		fmt.Fprintf(w, "empty string provided\n")
		return
	}
	if len(pat) == 0 {
		fmt.Fprintf(w, "empty pattern provided\n")
		return
	}

	for i := 0; i < len(txt)-1; {
		for j := 0; j <= len(pat)-1; {
			if txt[i] == pat[j] {
				if j == len(pat)-1 {
					ind = fmt.Sprintf("pattern starting index is %d", i-j)
					res = append(res, ind)
				}
				j++
				i++
			} else if txt[i] != pat[j] {
				break
			}
		}
		i++
	}
	if len(res) != 0 {
		fmt.Fprintf(w, "%v\n", res)
	} else {
		fmt.Fprintf(w, "pattern not found in string\n")
	}
}

func main() {

}
