package medium

import "sort"

// 2580. Count Ways to Group Overlapping Ranges
func countWays(ranges [][]int) int {
	MOD := 1000000000 + 7
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	n := len(ranges)
	curTail := -1
	num := 0
	for i := 0; i < n; i++ {
		l, r := ranges[i][0], ranges[i][1]
		if l > curTail {
			num++
			curTail = r
		} else {
			if curTail < r {
				curTail = r
			}
		}
	}
	ret := 1
	for num > 0 {
		ret = (ret * 2) % MOD
		num--
	}
	return ret
}
