package medium

import (
	"container/heap"
)

// 2462. Total Cost to Hire K Workers
func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	h := &Heap{}
	left, right := candidates-1, n-candidates
	if left+1 < right {
		for i := 0; i <= left; i++ {
			heap.Push(h, []int{costs[i], i})
		}
		for i := right; i < n; i++ {
			heap.Push(h, []int{costs[i], i})
		}
	} else {
		for i := 0; i < n; i++ {
			heap.Push(h, []int{costs[i], i})
		}
	}
	ans := int64(0)
	for i := 0; i < k; i++ {
		p := heap.Pop(h).([]int)
		cost, id := p[0], p[1]
		ans += int64(cost)
		if left+1 < right {
			if id <= left {
				left++
				heap.Push(h, []int{costs[left], left})
			} else {
				right--
				heap.Push(h, []int{costs[right], right})
			}
		}
	}
	return ans
}

type Heap [][]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}

//这两个Swap的实现都是一样的.
//func (h Heap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	ret := old[n-1]
	*h = old[0 : n-1]
	return ret
}
