package flyweight

// Flyweight 抽象享元角色
type Flyweight interface {
	data() string
}

// concreteFlyweight 具体享元角色
type concreteFlyweight struct {
	content string
}

func (cf *concreteFlyweight) data() string {
	return cf.content
}

func NewConcreteFlyweight(content string) *concreteFlyweight {
	content = "sharable:" + content
	return &concreteFlyweight{content}
}

// FlyweightFactory 享元工厂
type FlyweightFactory map[string]Flyweight

func (f FlyweightFactory) get(name string) Flyweight {
	flyweight := f[name]
	if flyweight == nil {
		f[name] = NewConcreteFlyweight(name)
		flyweight = f[name]
	}
	return flyweight
}

// flyweightFactory 具体的享元工厂
var flyweightFactory FlyweightFactory

func GetFlyweightFactory() FlyweightFactory {
	if flyweightFactory == nil {
		flyweightFactory = make(FlyweightFactory)
	}
	return flyweightFactory
}

// 非享元角色
type unsharableFlyweight struct {
	f Flyweight
}

func (u *unsharableFlyweight) Display() string {
	return "unsharable:" + u.f.data()
}

func NewUnsharableFlyweight() *unsharableFlyweight {
	f := GetFlyweightFactory().get("a")
	return &unsharableFlyweight{f: f}
}
