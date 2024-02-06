package array

// 1690. Stone Game VII
// 这个程序，对go写法没啥提高
func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + stones[i]
	}
	for i := n - 2; i >= 0; i-- { //这个两层循环还不错
		for j := i + 1; j < n; j++ {
			// 得分规则是，拿了一头，得剩下的分
			dp[i][j] = max(sum[j+1]-sum[i+1]-dp[i+1][j], sum[j]-sum[i]-dp[i][j-1])
		}
	}
	return dp[0][n-1]
}
