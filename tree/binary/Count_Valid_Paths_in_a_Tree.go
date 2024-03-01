package binary

// 2867. Count Valid Paths in a Tree
func countPaths(n int, edges [][]int) int64 {
	G := make([][]int, n+1)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		G[i] = append(G[i], j)
		G[j] = append(G[j], i)
	}
	var dfs func(int, int)
	var seen []int
	dfs = func(i int, pre int) {
		seen = append(seen, i)
		for _, j := range G[i] {
			if j != pre && !is_prime[j] {
				dfs(j, i)
			}
		}
	}
	res := int64(0)
	count := make([]int64, n+1)
	for i := 1; i < n; i++ {
		if !is_prime[i] {
			continue
		}
		cur := int64(0)
		for _, j := range G[i] {
			if is_prime[j] {
				continue
			}
			// j点是非质数点，j是i(质数点的一个临近的点)
			if count[j] == 0 {
				// 这就传参了，niubility
				seen = []int{} //注意这里是赋值，而不能是“声明+赋值”
				dfs(j, 0)      //这一下就把与j相关的这一大坨就访问完毕了，而且放到了seen里面，而且seen中的这些节点的count都相同的
				cnt := int64(len(seen))
				for _, k := range seen {
					count[k] = cnt
				}
			}
			res += count[j] * cur
			cur += count[j]
		}
		res += cur
	}
	return res
}

const N = 100001

var is_prime [N]bool

func init() {
	for i := 0; i < N; i++ {
		is_prime[i] = true
	}
	is_prime[1] = false
	for i := 2; i*i < N; i++ {
		if is_prime[i] {
			for j := i * i; j < N; j += i {
				is_prime[j] = false
			}
		}
	}
}
