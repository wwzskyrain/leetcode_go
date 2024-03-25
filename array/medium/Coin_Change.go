package medium

// 322. Coin Change
// 01背包问题：不用记，只需要理解其递归公式即可
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	MAX := 10000 + 1
	for i := range dp {
		dp[i] = MAX
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		min := dp[i]
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]+1 < min {
				min = dp[i-coin] + 1
			}
		}
		dp[i] = min
	}
	if dp[amount] == MAX {
		return -1
	}
	return dp[amount]
}
