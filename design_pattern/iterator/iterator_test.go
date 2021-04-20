package iterator

import "testing"

func TestIterator(t *testing.T) {
	l := newLists()

	eles := []interface{}{
		"sss", 123, true, false, "aaa", 0,
	}
	for i := range eles {
		l.add(eles[i])
	}

	iter := l.iterator()
	for i := 0; i < len(eles); i++ {
		ele := iter.next()
		if ele != eles[i] {
			t.Errorf("returns ele %v, expected %v\n", ele, eles[i])
		}
	}
}
