package hard

import "sort"

func minOperations(nums []int) int {
	sort.Ints(nums)
	var n = len(nums)
	var m = 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[m] = nums[i] //原地去重
			m++
		}
	}
	var ans = 0
	var left = 0
	for i := 0; i < m; i++ {
		for nums[left] < nums[i]-n+1 {
			left++
		}
		newAns := i - left + 1
		if ans < newAns {
			ans = newAns
		}
	}
	return n - ans
}
