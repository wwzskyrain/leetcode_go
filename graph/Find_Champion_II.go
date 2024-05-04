package graph

// 2924. Find Champion II
func findChampion(n int, edges [][]int) int {
	indgree := make([]int, n)
	for _, e := range edges {
		_, y := e[0], e[1]
		indgree[y]++
	}
	ans := -1
	for idx, in := range indgree {
		if in == 0 {
			if ans == -1 {
				ans = idx
			} else {
				return -1
			}
		}
	}
	return ans
}
