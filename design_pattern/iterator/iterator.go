package iterator

// Aggregate 抽象聚合
type Aggregate interface {
	add(i interface{})
	remove(idx int)
	iterator() Iterator
}

// Iterator 抽象迭代器
type Iterator interface {
	first()
	hasNext() bool
	next() interface{}
}

// lists 具体聚合
type lists struct {
	l []interface{}
}

func (l *lists) add(i interface{}) {
	l.l = append(l.l, i)
}

func (l *lists) remove(idx int) {
	copy(l.l[:idx], l.l[idx+1:])
	l.l = l.l[:len(l.l)-1]
}

func (l *lists) iterator() Iterator {
	return newListIterator(l)
}

func newLists() *lists {
	return &lists{l: make([]interface{}, 0)}
}

// listIterator 具体迭代器
type listIterator struct {
	l     *lists
	index int
}

func (lt *listIterator) first() {
	lt.index = 0
}

func (lt *listIterator) hasNext() bool {
	return lt.index < len(lt.l.l)
}

func (lt *listIterator) next() interface{} {
	if lt.hasNext() {
		ele := lt.l.l[lt.index]
		lt.index++
		return ele
	}
	return nil
}

func newListIterator(l *lists) *listIterator {
	return &listIterator{l, 0}
}
