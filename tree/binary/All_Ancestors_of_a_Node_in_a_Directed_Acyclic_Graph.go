package binary

import "sort"

// 2192. All Ancestors of a Node in a Directed Acyclic Graph
func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
	}
	in := make([]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		g[from] = append(g[from], to)
		in[to]++
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	parents := make([]map[int]bool, n)
	for i := range parents {
		parents[i] = map[int]bool{}
	}
	for len(queue) != 0 {
		root := queue[0]
		queue = queue[1:]
		rootParents := parents[root]
		for _, child := range g[root] {
			childParent := parents[child]
			for p := range rootParents {
				childParent[p] = true
			}
			childParent[root] = true //
			in[child]--
			if in[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	ans := make([][]int, n)
	for i := range ans {
		pMap := parents[i]
		for k, _ := range pMap {
			ans[i] = append(ans[i], k)
		}
		sort.Ints(ans[i])
	}
	return ans
}
