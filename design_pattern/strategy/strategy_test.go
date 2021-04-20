package strategy

import "testing"

func TestStrategy(t *testing.T) {
	c := newCash()
	b := newBank()

	pc := newPayment("tom", 100, c)
	pb := newPayment("tom", 100, b)

	if pc.pay() != "Pay 100 to tom by cash" {
		t.Error("wrong cash implementor")
	}

	if pb.pay() != "Pay 100 to tom by bank account" {
		t.Error("wrong bank implementor")
	}
}
