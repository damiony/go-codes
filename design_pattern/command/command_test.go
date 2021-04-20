package command

import "testing"

func TestCommand(t *testing.T) {
	op1 := newRebootCommand()
	op2 := newStartCommand()

	inv := newInvoker(op1, op2)
	if inv.excute() != "reboot:start" {
		t.Error("command error")
	}
}
