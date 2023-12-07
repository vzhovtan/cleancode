//Package binsearch implements binary search for []int arrays
package binsearch

func search(arr []int, lo, hi, n int) bool {
	for lo < hi {
		m := lo + (hi-lo)/2
		v := arr[m]
		if v == n {
			return true
		} else if v > n {
			hi = m
			continue
		} else {
			lo = m + 1
			continue
		}
	}
	return false
}

func main() {
}
