package tree

// 310. Minimum Height Trees
// 这个题目还是比较有意思的。这是一个图的拓扑排序的题目。
// 1.图的邻接矩阵表示法；
// 2.图的度；
// 3.就是拓扑排序本身了
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		degree[a]++
		degree[b]++
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var q []int
	for i, d := range degree {
		if d == 1 {
			q = append(q, i)
		}
	}
	var ans []int
	for len(q) > 0 {
		ans = []int{} //又是一轮赋值，如果这里的q中的值其入度都是1，可以确认，要么一个节点要么两个。
		// 该题目想找的也就是最长那条链路的中间节点（可一个可两个）。所以，拓扑排序也能解决寻找最长路径的问题吗？
		for i := len(q); i > 0; i-- {
			//一个一个的出队吧。
			a := q[0]
			q = q[1:] //这种就是go中的队列的使用方法
			ans = append(ans, a)
			//	把a删除掉
			for _, b := range g[a] {
				degree[b]--
				if degree[b] == 1 {
					q = append(q, b)
				}
			}
		}
	}
	return ans
}
