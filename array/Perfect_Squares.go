package array

import "math"

// 279. Perfect Squares
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i
	}
	//
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for j := i * i; j < len(dp); j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}
