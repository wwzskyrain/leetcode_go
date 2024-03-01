package binary

// 2673. Make Costs of Paths Equal in a Binary Tree
func minIncrements(n int, cost []int) int {
	res := 0
	for i := n - 2; i > 0; i -= 2 {
		res += abs(cost[i] - cost[i+1])
		cost[i/2] += max(cost[i], cost[i+1])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
