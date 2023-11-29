//Given a string S, find the length of the longest substring T that contains at most two
//distinct characters.
//For example,
//Given S = “eceba”,
//T is "ece" which its length is 3.
package longest2chars

import "fmt"

func long2c(inp string) int {
	max := 0
	cnt := map[rune]int{}
	numDist := 0
	i := 0
	for j, v := range inp {
		if cnt[v] == 0 {
			numDist++
		}
		cnt[v]++
		for numDist > 2 {
			cnt[rune(inp[i])]--
			if cnt[rune(inp[i])] == 0 {
				numDist--
				i++
			}
			if max < j-i {
				max = j - i
			}
		}
	}
	return max
}

func main() {
	fmt.Println(long2c("abaac"))
}
