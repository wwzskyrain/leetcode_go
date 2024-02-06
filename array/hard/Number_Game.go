package hard

import (
	"container/heap"
	"sort"
)

//LCP 24. 数字游戏

func numsGame(nums []int) []int {
	const mod = 1_000_000_007
	ans := make([]int, len(nums))
	left := hp{}  // 维护较小的一半，大根堆（小根堆取负号）
	right := hp{} // 维护较大的一半，小根堆
	for i, b := range nums {
		b -= i
		if i%2 == 0 {
			heap.Push(&right, -left.pushPop(-b))
			x := right.IntSlice[0] // 中位数
			// 原本要减去 left.sum，但由于 left 所有元素都取负号，所以负负得正改为加法
			ans[i] = (right.sum - x + left.sum) % mod
		} else {
			heap.Push(&left, -right.pushPop(b))
			ans[i] = (right.sum + left.sum) % mod
		}
	}
	return ans
}

type hp struct {
	sort.IntSlice     // 继承 Len, Less, Swap
	sum           int // 堆中元素之和
}

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)); h.sum += v.(int) }
func (hp) Pop() (_ any)  { return } // 没用到，无需实现

// pushPop 先把 v 入堆，然后弹出并返回堆顶
// 如果 v <= 堆顶，则直接返回 v
func (h *hp) pushPop(v int) int {
	if h.Len() > 0 && v > h.IntSlice[0] {
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}
