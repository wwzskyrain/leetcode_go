package medium

// 377. Combination Sum IV
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(dp); i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}
	return dp[target]
}
