package template

// templater 抽象类的接口 非必须
type templater interface {
	games() string
	music() string
	Print() string
}

// subClasser 子类接口
type subClasser interface {
	games() string
	music() string
}

// templateImpl 抽象类的具体实现
type templateImpl struct {
	subClasser
}

func (t *templateImpl) Print() string { return t.subClasser.games() + ":" + t.subClasser.music() }

func (t *templateImpl) games() string { return "" }

func (t *templateImpl) music() string { return "" }

func NewTemplateImpl(sub subClasser) *templateImpl {
	return &templateImpl{sub}
}

// phone 子类 手机
type phone struct {
	templater
}

func (p *phone) games() string {
	return "phone games"
}

func (p *phone) music() string {
	return "phone music"
}

func NewPhone() *phone {
	p := &phone{}
	t := NewTemplateImpl(p)
	p.templater = t
	return p
}

// computer 子类 电脑
type computer struct {
	templater
}

func (c *computer) games() string {
	return "computer games"
}

func (c *computer) music() string {
	return "computer music"
}

func NewComputer() *computer {
	c := &computer{}
	p := NewTemplateImpl(c)
	c.templater = p
	return c
}
