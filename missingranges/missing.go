package missingranges

import (
	"fmt"
	"strings"
)

func findMissing(inp []int) string {
	var b strings.Builder

	if len(inp) == 0 {
		b.WriteString("{0->9}")
		return b.String()
	}

	b.WriteString("{")
	for i := 1; i < len(inp); i++ {
		if inp[i]-inp[i-1] >= 2 {
			rng := getRange(inp[i-1], inp[i])
			if len(b.String()) > 1 {
				b.WriteString(", ")
			}
			b.WriteString(rng)
		}
	}
	if len(b.String()) == 0 {
		return "{}"
	} else {
		b.WriteString("}")
		return b.String()
	}
}

func getRange(start, end int) string {
	if (end - start) == 2 {
		return fmt.Sprintf("%d", start+1)
	} else {
		return fmt.Sprintf("%d -> %d", start+1, end-1)
	}
}

func main() {
	in := []int{0, 1, 3, 5, 66, 89}
	fmt.Println(findMissing(in))
}
