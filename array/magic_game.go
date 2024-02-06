package array

import (
	"container/heap"
	"sort"
)

func magicTower(nums []int) (ans int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < 0 {
		return -1
	}
	hp := 1
	h := &minHeap{}
	ans = 0
	for _, num := range nums {
		if num < 0 {
			heap.Push(h, num)
		}
		hp += num
		if hp < 1 {
			// pop
			pop := heap.Pop(h)
			hp -= pop.(int)
			ans++
		}
	}
	return ans
}

type minHeap struct {
	sort.IntSlice
}

func (h *minHeap) Push(v any) {
	//放到最后一个就行，会up调整的
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *minHeap) Pop() any {
	//把最后一个元素删除并返回，放心，调用该方法之前已经把堆顶元素放到了这最后一个位置
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
