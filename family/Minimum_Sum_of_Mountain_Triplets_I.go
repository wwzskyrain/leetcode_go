package family

// 2908. Minimum Sum of Mountain Triplets I
func minimumSum(nums []int) int {
	var n = len(nums)
	var left = make([]int, n)
	left[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < left[i-1] {
			left[i] = nums[i]
		} else {
			left[i] = left[i-1]
		}
	}
	var rightMin = nums[n-1]
	var ret = 10000
	for i := n - 2; i > 0; i-- {
		if nums[i] > left[i-1] && nums[i] > rightMin {
			var sum = rightMin + nums[i] + left[i-1]
			if sum < ret {
				ret = sum
			}
		}
		if rightMin > nums[i] {
			rightMin = nums[i]
		}
	}
	if ret < 10000 {
		return ret
	}
	return -1
}
