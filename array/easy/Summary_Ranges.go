package easy

import "strconv"

// 228. Summary Ranges
// 又在ai的帮助下，攻克一题。
func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i] == nums[i-1]+1; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return ans
}
