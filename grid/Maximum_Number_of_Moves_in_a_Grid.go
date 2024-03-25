package grid

// 2684. Maximum Number of Moves in a Grid
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		ans = max(ans, j)
		if ans == n-1 {
			return
		}
		for k := max(i-1, 0); k < min(i+2, m); k++ {
			if grid[k][j+1] > grid[i][j] {
				dfs(k, j+1)
			}
		}
		grid[i][j] = 0 //不再访问了
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
	}
	return ans
}
