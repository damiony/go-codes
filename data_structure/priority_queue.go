package data_structure

import (
	"container/heap"
	"fmt"
)

// 优先级队列
type PriorityNode struct {
	data     interface{}
	priority int // 优先级
	index    int // 进行更新时需要用到
}

type PriorityQueue []*PriorityNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = pq[j].index, pq[i].index
}

func (pq *PriorityQueue) Push(node interface{}) {
	n := node.(*PriorityNode)
	n.index = len(*pq)
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	node := (*pq)[n-1]
	node.index = -1
	*pq = (*pq)[:n-1]
	return node
}

func (pq *PriorityQueue) Update(node *PriorityNode, data interface{}, priority int) {
	node.data = data
	node.priority = priority
	heap.Fix(pq, node.index)
}

func PriorityQueueExamples() {
	priority := map[string]int{
		"apple":  15,
		"banana": 10,
		"orange": 14,
	}
	// 构造优先级队列
	pq := PriorityQueue{}
	index := 0
	for k, v := range priority {
		pq = append(pq, &PriorityNode{
			data:     k,
			priority: v,
			index:    index,
		})
		index++
	}
	// 堆化
	heap.Init(&pq)
	// 添加元素
	node := PriorityNode{
		data:     "watermelon",
		priority: 20,
	}
	heap.Push(&pq, &node)
	pq.Update(&node, node.data, 9)
	// 按照优先级输出
	for len(pq) > 0 {
		fmt.Println(heap.Pop(&pq))
	}
}
