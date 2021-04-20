package visitor

import "testing"

func TestVisitor(t *testing.T) {
	o := newObjectStructure()

	e1 := newConcreteElementA()
	e2 := newConcreteElementB()
	o.add(e1)
	o.add(e2)

	v1 := newConcreteVisitorA()
	v2 := newConcreteVisitorB()
	if o.accept(v1) != "concreteVisitorA:concreteElementA" {
		t.Error("visitorA")
	}
	if o.accept(v2) != "concreteVisitorB:concreteElementB" {
		t.Error("visitorB")
	}
}
