package medium

import "fmt"

// 299. Bulls and Cows
func getHint(secret string, guess string) string {
	n := len(secret)
	bull := 0
	s1, s2 := make([]int, 10), make([]int, 10)
	for i := 0; i < n; i++ {
		i1, i2 := secret[i]-'0', guess[i]-'0'
		if i1 == i2 {
			bull++
		} else {
			s1[i1]++
			s2[i2]++
		}
	}
	cow := 0
	for i := 0; i < 10; i++ {
		cow += min(s1[i], s2[i])
	}
	return fmt.Sprintf("%dA%dB", bull, cow)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
