package facade

// Facade 外观者接口
type Facade interface {
	method() string
}

// facadeImpl 具体的外观者实现
type facadeImpl struct {
	a SubSystemA
	b SubSystemB
}

func (f *facadeImpl) method() string {
	return f.a.method() + f.b.method()
}

// NewFacade 外观者工厂
func NewFacade() Facade {
	return &facadeImpl{
		a: NewSubSystemA(),
		b: NewSubSystemB(),
	}
}

// SubSystemA 子系统A接口
type SubSystemA interface {
	method() string
}

// SubSystemB 子系统B接口
type SubSystemB interface {
	method() string
}

// subSystemA 子系统A具体实现
type subSystemA struct{}

func (a *subSystemA) method() string {
	return "subsystemA\n"
}

func NewSubSystemA() *subSystemA {
	return &subSystemA{}
}

// subSystemB 子系统B具体实现
type subSystemB struct{}

func (b *subSystemB) method() string {
	return "subsystemB\n"
}

func NewSubSystemB() *subSystemB {
	return &subSystemB{}
}
