package medium

import "sort"

// 56. Merge Intervals
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	n := len(intervals)
	tail := -1
	var start int
	for i := 0; i < n; i++ {
		intverval := intervals[i]
		if tail < intverval[0] {
			if tail >= 0 {
				ret = append(ret, []int{start, tail})
			}
			start = intverval[0]
			tail = intverval[1]
		} else {
			if tail < intverval[1] {
				tail = intverval[1]
			}
		}
	}
	if tail >= intervals[n-1][1] {
		ret = append(ret, []int{start, tail})
	}
	return ret
}
