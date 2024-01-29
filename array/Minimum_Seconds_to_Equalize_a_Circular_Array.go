package array

// 2808. Minimum Seconds to Equalize a Circular Array
// 1. 找相邻元素e之间的最大距离d，则把该nums变成该元素e需要d/2
// 2. 比较最小的d/2即可
func minimumSeconds(nums []int) int {
	mp := make(map[int][]int)
	n := len(nums)
	res := n
	for i, num := range nums {
		mp[num] = append(mp[num], i)
	}
	for _, pos := range mp {
		mx := pos[0] + n - pos[len(pos)-1]
		for i := 1; i < len(pos); i++ {
			mx = max(mx, pos[i]-pos[i-1])
		}
		res = min(res, mx/2)
	}
	return res
}

func max(a, b int) (larger int) {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) (less int) {
	if a > b {
		return b
	}
	return a
}
