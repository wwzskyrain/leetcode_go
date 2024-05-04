package tree

// 1766. Tree of Coprimes

func getCoprimes(nums []int, edges [][]int) []int {
	MAX := 51
	n := len(nums)
	g := make([][]int, n)
	gcds := make([][]int, MAX)
	tmp := make([][]int, MAX)
	dep := make([]int, n)
	ans := make([]int, n)

	// 初始化
	for i := 0; i < 51; i++ {
		gcds[i] = []int{}
		tmp[i] = []int{}
	}
	for i := 0; i < n; i++ {
		g[i] = []int{}
		ans[i], dep[i] = -1, -1
	}

	// 初始化gcds
	for i := 1; i < MAX; i++ {
		for j := 1; j < MAX; j++ {
			if gcd(i, j) == 1 {
				gcds[i] = append(gcds[i], j)
			}
		}
	}
	//初始化g
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(x, depth int)

	dfs = func(x, depth int) {
		dep[x] = depth
		for _, val := range gcds[nums[x]] {
			//tmp[val] 存放的是node节点序号，从0到n-1，而且是先进后出的栈，栈顶就是最深的层次
			pathNodeValLen := len(tmp[val])
			if pathNodeValLen == 0 {
				continue
			}
			lastNodeVal := tmp[val][pathNodeValLen-1]
			if ans[x] == -1 || dep[lastNodeVal] > dep[ans[x]] {
				ans[x] = lastNodeVal
			}
		}
		// tep存放的是dfs遍历过程中-访问中的节点的val和其节点序号x的映射
		tmp[nums[x]] = append(tmp[nums[x]], x)
		for _, child := range g[x] {
			if dep[child] == -1 { //被访问过的节点的深度肯定不为-1
				dfs(child, depth+1)
			}
		}
		//pop()
		tmp[nums[x]] = tmp[nums[x]][:len(tmp[nums[x]])-1]
	}
	dfs(0, 1)
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
