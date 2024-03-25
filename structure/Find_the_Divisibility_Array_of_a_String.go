package structure

// 2575. Find the Divisibility Array of a String
func divisibilityArray(word string, m int) []int {
	cur := 0
	ret := make([]int, 0)
	for _, c := range word {
		cur = (cur*10 + int(c-'0')) % m
		if cur == 0 {
			ret = append(ret, 1)
		} else {
			ret = append(ret, 0)
		}
	}
	return ret
}
