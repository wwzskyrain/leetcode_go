package medium

import "sort"

// 826. Most Profit Assigning Work
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	jobs := make([][2]int, len(difficulty))
	for i, diff := range difficulty {
		jobs[i] = [2]int{diff, profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	sort.Ints(worker)
	res, i, best := 0, 0, 0
	for _, w := range worker {
		for i < len(difficulty) && w > jobs[i][0] {
			best = max(best, jobs[i][1])
			i++
		}
		res += best
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
