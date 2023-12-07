package longest2chars

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
}
