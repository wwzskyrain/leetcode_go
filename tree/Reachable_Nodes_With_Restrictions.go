package tree

// 2368. Reachable Nodes With Restrictions
func reachableNodes(n int, edges [][]int, restricted []int) int {
	g := make([][]int, n) //用二维切片来当邻接矩阵，注意，这里一定要初始化切片长度，因为赋值是随机的
	cnt := 0
	for _, edge := range edges {
		v1, v2 := edge[0], edge[1]
		g[v1] = append(g[v1], v2)
		g[v2] = append(g[v2], v1)
	}
	r := make(map[int]int) //map这里就不用考虑长度了，
	for _, node := range restricted {
		r[node] = 1
	}
	var dfs func(g [][]int, root int, pre int, set map[int]int)
	dfs = func(g [][]int, root int, pre int, set map[int]int) {
		if _, ok := set[root]; ok {
			return
		}
		cnt++
		for _, next := range g[root] {
			if pre != next {
				dfs(g, next, root, set)
			}
		}
	}

	dfs(g, 0, -1, r)
	return cnt
}
