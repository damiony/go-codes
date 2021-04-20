package component

import "fmt"

// Component抽象构件
type Component interface {
	add(c Component)
	remove(c Component)
	getChild(name string) Component
	setname(name string)
	getname() string
	print() string
}

// componentBase 具体构件的基类
type componentBase struct {
	name string
}

func (base *componentBase) add(child Component) {}

func (base *componentBase) remove(child Component) {}

func (base *componentBase) getChild(name string) Component { return nil }

func (base *componentBase) setname(name string) { base.name = name }

func (base *componentBase) getname() string { return base.name }

func (base *componentBase) print() string { return "" }

// Leaf 树叶构件
type Leaf struct {
	componentBase
}

func (l *Leaf) print() string {
	return fmt.Sprintln("leaf:", l.name)
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

// Composite 树枝构件
type Composite struct {
	componentBase
	children []Component
}

func (c *Composite) print() string {
	s := fmt.Sprintln("composite:", c.name)
	for _, child := range c.children {
		s += child.print()
	}
	return s
}

func (c *Composite) add(child Component) {
	c.children = append(c.children, child)
}

func (c *Composite) remove(child Component) {
	for i, v := range c.children {
		if v.getname() == child.getname() {
			copy(c.children[:i], c.children[i+1:])
			c.children = c.children[:len(c.children)-1]
			break
		}
	}
}

func (c *Composite) getChild(name string) Component {
	for _, v := range c.children {
		if v.getname() == name {
			return v
		}
	}
	return nil
}

func NewComposite() *Composite {
	return &Composite{
		children: make([]Component, 0),
	}
}
