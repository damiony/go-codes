package data_structure

type queueNode struct {
	data interface{}
	next *queueNode
}

// 链式队列
type ChainedQueue struct {
	length int
	head   *queueNode
	tail   *queueNode
}

func NewChainedQ() (q *ChainedQueue) {
	return
}

func (q *ChainedQueue) Push(x interface{}) {
	node := &queueNode{
		data: x,
	}

	q.length++
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = q.tail.next
}

func (q *ChainedQueue) Pop() interface{} {
	node := q.head
	q.head = q.head.next
	q.length--
	return node.data
}

func (q *ChainedQueue) Top() interface{} {
	return q.head.data
}

func (q *ChainedQueue) Len() int {
	return q.length
}

func (q *ChainedQueue) IsEmpty() bool {
	return q.length == 0
}
