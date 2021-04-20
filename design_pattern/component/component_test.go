package component

import "testing"

func TestComponent(t *testing.T) {
	// 树枝构件
	c0 := NewComposite()
	c0.setname("c0")
	c1 := NewComposite()
	c1.setname("c1")
	// 树叶构件
	l0 := NewLeaf()
	l0.setname("l0")
	l1 := NewLeaf()
	l1.setname("l1")
	l2 := NewLeaf()
	l2.setname("l2")

	c0.add(l0)
	c0.add(c1)
	c1.add(l1)
	c1.add(l2)

	s := c0.print()
	if s != "composite: c0\nleaf: l0\ncomposite: c1\nleaf: l1\nleaf: l2\n" {
		t.Errorf("wrong: %v", s)
	}
}
