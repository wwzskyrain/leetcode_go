package medium

import (
	"sort"
)

// 15. 三数之和
// 温习一下思路；而go，也就是排序sort.Ints(nums)的使用吧
func threeSum(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	for first := 0; first+2 < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		target := -nums[first]
		for second := first + 1; second+1 < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ret = append(ret, []int{nums[first], nums[second], nums[third]})
			}
			// 执行到这里，nums[second] + nums[third] < target了，也就是second要往前挪一步了。
		}
	}
	return ret
}
