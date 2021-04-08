package data_structure

// 线性队列
type LinearQueue struct {
	nums []int
}

func NewLinearQ() (l LinearQueue) {
	return
}

func (l *LinearQueue) Push(x int) {
	l.nums = append(l.nums, x)
}

func (l *LinearQueue) Pop() int {
	x := l.nums[0]
	l.nums = l.nums[1:]
	return x
}

func (l *LinearQueue) Top() int {
	return l.nums[0]
}

func (l *LinearQueue) Len() int {
	return len(l.nums)
}

func (l *LinearQueue) IsEmpty() bool {
	return len(l.nums) == 0
}
