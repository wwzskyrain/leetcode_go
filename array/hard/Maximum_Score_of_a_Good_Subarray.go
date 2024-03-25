package hard

// 1793. Maximum Score of a Good Subarray
// 这个题目很不错。
//1.要知道为什么这样的遍历可以求出最大值
//	因为就两个因素，高和宽，高是从大到小，宽是从小到大，说到这里已经无需多言了吧。

// 2.要明白对i的更新：
//
//	更新的策略：取当前左右两边更小值的较大值
//	因为在高从大到小的过程，不能跳过这个更大值呀
func maximumScore(nums []int, k int) int {
	ans := 0
	n := len(nums)
	left, right := k-1, k+1
	for i := nums[k]; ; { // 这个i表示当前的最小值，其实有很多控制
		for left >= 0 && nums[left] >= i {
			left--
		}
		for right < n && nums[right] >= i {
			right++
		}
		ans = max(ans, i*(right-left-1))
		if left == -1 && right == n {
			break
		}
		//
		if left >= 0 && right < n {
			i = max(nums[left], nums[right])
		} else if left >= 0 {
			i = nums[left]
		} else {
			i = nums[right]
		}
	}
	return ans
}
