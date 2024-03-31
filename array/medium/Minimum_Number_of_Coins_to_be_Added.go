package medium

import "sort"

// 2952. Minimum Number of Coins to be Added
func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	var n = len(coins)
	var next = 1
	var i = 0
	var ret = 0
	for next <= target {
		if i < n && coins[i] <= next {
			next += coins[i]
			i++
		} else {
			next *= 2
			ret++
		}
	}
	return ret
}
