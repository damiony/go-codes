package strategy

import "fmt"

// payment 环境类
type payment struct {
	name    string
	money   int
	trategy trategyer
}

func (p *payment) pay() string {
	return p.trategy.pay(p.name, p.money)
}

func newPayment(name string, money int, t trategyer) *payment {
	return &payment{name, money, t}
}

// trategyer 抽象策略
type trategyer interface {
	pay(name string, money int) string
}

// cash 具体策略
type cash struct{}

func (c *cash) pay(name string, money int) string {
	return fmt.Sprintf("Pay %d to %s by cash", money, name)
}

func newCash() *cash {
	return &cash{}
}

// bank 具体策略
type bank struct{}

func (b *bank) pay(name string, money int) string {
	return fmt.Sprintf("Pay %d to %s by bank account", money, name)
}

func newBank() *bank {
	return &bank{}
}
