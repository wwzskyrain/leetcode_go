package medium

// 2369. Check if There is a Valid Partition For The Array
func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, len(nums)+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		if i >= 2 {
			dp[i] = dp[i-2] && towSame(nums[i-1], nums[i-2])
		}
		if i >= 3 {
			// 我靠，原来错在传参上，顺序反了
			dp[i] = dp[i] || (dp[i-3] && (threeValid(nums[i-3], nums[i-2], nums[i-1])))
		}
	}
	return dp[n]
}

func towSame(n1, n2 int) bool {
	return n1 == n2
}

func threeValid(n1, n2, n3 int) bool {
	return (n1 == n2 && n2 == n3) || (n2-n1 == 1 && n3-n2 == 1)
}

// 递归法，肯定超时呀
func validPartition1(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	if len(nums) == 1 {
		return false
	}

	if nums[0] == nums[1] {
		if validPartition(nums[2:]) {
			return true
		}
		if len(nums) > 2 && nums[1] == nums[2] {
			if validPartition(nums[3:]) {
				return true
			}
		}
	}
	if len(nums) > 2 && nums[0]+1 == nums[1] && nums[1]+1 == nums[2] {
		if validPartition(nums[3:]) {
			return true
		}
	}
	return false
}
