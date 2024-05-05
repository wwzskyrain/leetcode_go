package easy

// 1652. Defuse the Bomb
// wo服了，几乎都是生成的代码.
func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}
	code = append(code, code...)
	l, r := 1, k
	if k < 0 {
		l, r = n+k, n-1
	}
	sum := 0
	for i := l; i <= r; i++ {
		sum += code[i]
	}
	for i := 0; i < n; i++ {
		ans[i] = sum
		sum -= code[l]
		l++
		r++
		sum += code[r]
	}
	return ans
}
