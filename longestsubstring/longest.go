//Given a string, find the length of the longest substring without repeating characters. For
//example, the longest substring without repeating letters for “abcabcbb” is “abc”, which
//the length is 3. For “bbbbb” the longest substring is “b”, with the length of 1
package longestsubstring

import (
	"fmt"
)

func longest(inp string) int {
	hash := map[rune]int{}
	head := 0
	max := 0
	for tail, v := range inp {
		if hash[v] >= head {
			head = hash[v] + 1
		}
		hash[v] = tail
		if max < tail-head+1 {
			max = tail - head + 1
		}
	}
	return max
}

func main() {
	astr := "abcabcbb"
	fmt.Println(longest(astr))
}
