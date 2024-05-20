package hard

// 1542. Find Longest Awesome Substring
func longestAwesome(s string) int {
	const D = 10
	n := len(s)
	pos := [1 << D]int{}
	for i := range pos {
		pos[i] = n // has not found pre
	}
	pos[0] = -1
	pre := 0 // only need a var instead of arr
	ans := 0
	for i, c := range s {
		pre ^= 1 << (c - '0')
		for d := 0; d < D; d++ {
			ans = max(ans, i-pos[pre^(1<<d)]) // odd length, so need pre1 - pre2 = 2^n, n from 1 to D
		}
		ans = max(ans, i-pos[pre]) // even length
		if pos[pre] == n {         // first find pre index
			pos[pre] = i
		}
	}
	return ans
}
