package medium

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	for i := 1; i < len(items); i++ {
		items[i][1] = max(items[i][1], items[i-1][1])
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = query(items, q)
	}
	return res
}

func query(items [][]int, q int) int {
	l, r := 0, len(items)
	for l < r {
		mid := l + (r-l)/2
		if items[mid][0] > q {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 {
		// 此时所有物品价格均大于查询价格
		return 0
	} else {
		// 返回小于等于查询价格的物品的最大美丽值
		return items[l-1][1]
	}
}
