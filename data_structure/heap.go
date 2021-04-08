package data_structure

import (
	"container/heap"
	"fmt"
)

// 最小堆
type MyHeap []int

func (m MyHeap) Len() int { return len(m) }

func (m MyHeap) Less(i, j int) bool { return m[i] < m[j] }

func (m MyHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *MyHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MyHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}

func HeapExamples() {
	h := &MyHeap{}
	heap.Init(h)
	fmt.Println("init heap: ", h)

	heap.Push(h, 45)
	heap.Push(h, 9)
	heap.Push(h, 3)
	heap.Push(h, 7)
	heap.Push(h, 4)
	heap.Push(h, 6)
	fmt.Println("push numbers: ", h)

	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
}
