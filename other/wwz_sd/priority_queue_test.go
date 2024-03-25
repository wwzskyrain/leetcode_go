package wwz_sd

import (
	"fmt"
	"testing"
)

// 我终于把这个优先级队列写出来了
func TestPriorityQueue(t *testing.T) {

	var pairs []Pair
	pairs = append(pairs, Pair{1, 1})
	pairs = append(pairs, Pair{2, 2})
	pairs = append(pairs, Pair{3, 3})
	pairs = append(pairs, Pair{4, 4})
	for i, pair := range pairs {
		fmt.Printf("&pair=%p, &pairs[%d] = %p\n", &pair, i, &pairs[i])
	}
	p2 := pairs[2]
	addPair(&p2) //注意，这里p2变量，和pair[2]是不同的变量了，他们是同模板下的不同的实例，它们有不同的地址。
	// 这样传递，“赋值变量的地址”，对pair[2]没有任何作用
	fmt.Printf("p2 = %v\n", p2)

	p2Addr := &pairs[2]
	addPair(p2Addr)    //1.间接传地址：传地址变量（一个变量，其值是地址）
	addPair(&pairs[2]) //2.直接传地址。
	fmt.Printf("&pairs[2] = %v\n", &pairs[2])

}

// 这是地址入参
func addPair(p *Pair) {
	fmt.Printf("p-> %p\n", p)
	(*p).First = (*p).First * 2
	(*p).Second = (*p).Second * 2
}

func TestMyPriorityQueueOut(t *testing.T) {
	pq := BuildPriorityQueueOut()
	pq.PushPair(Pair{5, 5})
	pq.PushPair(Pair{3, 3})
	fmt.Printf("peek = %v \n", pq.Peek())
	pq.PushPair(Pair{2, 2})
	pq.PushPair(Pair{1, 1})
	fmt.Printf("peek = %v \n", pq.Peek())
	pq.PushPair(Pair{3, 3})
	fmt.Printf("peek = %v \n", pq.Peek())

	for !pq.IsEmpty() {
		fmt.Printf("pop = %v \n", pq.PopPair())
	}

}
