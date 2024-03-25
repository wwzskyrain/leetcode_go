package wwz_sd

import (
	"container/heap"
)

type Pair struct {
	First  int
	Second int
}

type PriorityQueue []Pair

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].First < q[j].First
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

func (q *PriorityQueue) FirstEle() Pair {
	if len(*q) > 0 {
		f := (*q)[0]
		return f
	}
	return Pair{}
}

func (this *PriorityQueueOut) IsEmpty() bool {
	return this.PQ.Len() == 0
}

type PriorityQueueOut struct {
	PQ *PriorityQueue
}

func (this PriorityQueueOut) PushPair(x any) {
	heap.Push(this.PQ, x)
}

func (this *PriorityQueueOut) PopPair() Pair {
	x := heap.Pop(this.PQ)
	return x.(Pair)
}

// Peek 当队列空的时候，返回的是一个空结构体
func (this *PriorityQueueOut) Peek() Pair {
	return this.PQ.FirstEle()
}

func BuildPriorityQueueOut() *PriorityQueueOut {
	return &PriorityQueueOut{
		PQ: &PriorityQueue{},
	}
}
