package main

import (
	"container/heap"
	"fmt"
)

//IntHeap是int类型的最小堆实现
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	//方法绑定在一个指针类型上，因为方法改变了切片长度而不仅仅是内容。
	*h = append(*h, x.(int))
}

//
func (h *IntHeap) Pop() interface{} {
	//方法绑定在一个指针类型上，因为方法改变了切片长度而不仅仅是内容。
	old := *h
	n := len(old)
	x := n - 1
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{8, 1, 2, 3, 5, 6, 7, 0}
	heap.Init(h)
	fmt.Printf("the init head is %v\n", *h)
	heap.Pop(h)
	heap.Fix(h, 0)
	fmt.Printf("the poped head is %v\n", *h)
	heap.Push(h, 0)
	heap.Fix(h, 0)
	fmt.Printf("the pushed head is %v\n", *h)
}
