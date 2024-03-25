package hard

import "container/heap"

// 2617. Minimum Number of Visited Cells in a Grid
func minimumVisitedCells(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 1
	row := make([]PriorityQueue, m)
	col := make([]PriorityQueue, n)
	update := func(x *int, y int) {
		if *x == -1 || y < *x {
			*x = y
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//rowQ := row[i]
			for row[i].Len() > 0 && row[i][0].second+grid[i][row[i][0].second] < j {
				heap.Pop(&row[i]) // 搞笑，还要用这种写法，把优先级队列当做入参。
			}
			if row[i].Len() > 0 {
				update(&dist[i][j], dist[i][row[i][0].second]+1) // +1是一步跳的意思
			}
			colQ := col[j]
			for colQ.Len() > 0 && colQ[0].second+grid[colQ[0].second][j] < i {
				heap.Pop(&col[j]) //这里已经要注意传参
			}
			if col[j].Len() > 0 {
				update(&dist[i][j], dist[col[j][0].second][j]+1)
			}
			cur := dist[i][j]
			if dist[i][j] != -1 {
				heap.Push(&row[i], Pair{cur, j})
				heap.Push(&col[j], Pair{cur, i})
			}
		}
	}
	return dist[m-1][n-1]
}

type Pair struct {
	first  int
	second int
}

type PriorityQueue []Pair

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q PriorityQueue) Len() int {
	return len(q)
}
func (q PriorityQueue) Less(i, j int) bool {
	return q[i].first < q[j].first
}
func (q *PriorityQueue) Push(x any) {
	*q = append(*q, x.(Pair))
}
func (q *PriorityQueue) Pop() any {
	n := len(*q)
	x := (*q)[n-1]
	*q = (*q)[:n-1]
	return x
}
