package visitor

// Visitor 抽象访问者
type Visitor interface {
	visit(e Element) string
}

// Element 抽象元素
type Element interface {
	accept(v Visitor) string
}

// ConcreteVisitorA 具体访问者A
type ConcreteVisitorA struct{}

func (ca *ConcreteVisitorA) visit(e Element) string {
	s := ""
	switch e.(type) {
	case *ConcreteElementA:
		s = "concreteVisitorA:concreteElementA"
	}
	return s
}

func newConcreteVisitorA() *ConcreteVisitorA {
	return &ConcreteVisitorA{}
}

// ConcreteVisitorB 具体访问者B
type ConcreteVisitorB struct{}

func (cb *ConcreteVisitorB) visit(e Element) string {
	s := ""
	switch e.(type) {
	case *ConcreteElementB:
		s = "concreteVisitorB:concreteElementB"
	}
	return s
}

func newConcreteVisitorB() *ConcreteVisitorB {
	return &ConcreteVisitorB{}
}

// ConcreteElementA 具体元素A
type ConcreteElementA struct{}

func (ca *ConcreteElementA) accept(v Visitor) string {
	return v.visit(ca)
}

func newConcreteElementA() *ConcreteElementA {
	return &ConcreteElementA{}
}

// ConcreteElementB 具体元素B
type ConcreteElementB struct{}

func (cb *ConcreteElementB) accept(v Visitor) string {
	return v.visit(cb)
}

func newConcreteElementB() *ConcreteElementB {
	return &ConcreteElementB{}
}

// ObjectStructure 对象结构
type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) add(e Element) {
	o.elements = append(o.elements, e)
}

func (o *ObjectStructure) remove(i int) {
	copy(o.elements[:i], o.elements[i+1:])
	o.elements = o.elements[:len(o.elements)-1]
}

func (o *ObjectStructure) accept(v Visitor) string {
	res := ""
	for _, ele := range o.elements {
		res += ele.accept(v)
	}
	return res
}

func newObjectStructure() *ObjectStructure {
	return &ObjectStructure{}
}
