package reversewords

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//naive approach, splitting the string and then looping over resulted slice from the end
func reverseNaive(w io.Writer, inp string) {
	var b strings.Builder
	alist := strings.Split(inp, " ")
	for i := len(alist) - 1; i >= 0; i-- {
		if len(b.String()) != 0 {
			b.WriteString(" ")
		}
		b.WriteString(alist[i])
	}
	fmt.Fprintf(w, "%s\n", b.String())
}

//looping over the string from the end,  adding each word to new string
func reverseOnePass(w io.Writer, inp string) {
	var b strings.Builder
	var j = len(inp)
	for i := j - 1; i >= 0; i-- {
		if string(inp[i]) == " " {
			j = i
		} else if i == 0 || string(inp[i-1]) == " " {
			if len(b.String()) != 0 {
				b.WriteString(" ")
			}
			b.WriteString(inp[i:j])
		}
	}
	fmt.Fprintf(w, "%s\n", b.String())
}

func main() {
	inp := "the sky is blue"
	reverseNaive(os.Stdout, inp)
	reverseOnePass(os.Stdout, inp)
}
