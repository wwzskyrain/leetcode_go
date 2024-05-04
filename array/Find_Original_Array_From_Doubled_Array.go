package array

import "sort"

// 2007. Find Original Array From Doubled Array
func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	count := make(map[int]int)
	ret := make([]int, 0)
	for _, v := range changed {
		count[v]++
	}
	for _, v := range changed {
		if count[v] == 0 {
			continue
		}
		count[v]--
		if count[v*2] == 0 {
			return []int{} // return empty slice
		}
		count[v*2]--
		ret = append(ret, v)
	}
	return ret
}
