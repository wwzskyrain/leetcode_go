package main

import (
	"container/heap"
	"math"
)

// 1976. Number of Ways to Arrive at Destination
func countPaths(n int, roads [][]int) int {
	const mod = int64(1e9 + 7)
	e := make([][]Edge, n)
	for _, road := range roads {
		x, y, t := road[0], road[1], road[2]
		e[x] = append(e[x], Edge{y, t})
		e[x] = append(e[y], Edge{x, t})
	}
	dist := make([]int64, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt64
	}
	ways := make([]int64, n)
	q := PriorityQueue{{0, 0}}
	heap.Init(&q)
	dist[0] = 0
	ways[0] = 1

	for len(q) > 0 {
		p := heap.Pop(&q).(Pair)
		t, u := p.first, p.second
		if t > dist[u] {
			continue
		}
		for _, edge := range e[u] {
			v, w := edge.to, edge.cost
			if t+int64(w) < dist[v] {
				dist[v] = t + int64(w)
				ways[v] = ways[u]
				heap.Push(&q, Pair{t + int64(w), int64(v)})
			} else if t+int64(w) == dist[v] {
				ways[v] = (ways[u] + ways[v]) % mod
			}
		}
	}
	return int(ways[n-1])
}

type Edge struct {
	to   int
	cost int
}

type Pair struct {
	first, second int64
}

type PriorityQueue []Pair

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].first < pq[j].first
}
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Pair))
}
func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}
