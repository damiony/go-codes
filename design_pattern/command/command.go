package command

// commander 抽象命令
type commander interface {
	excute() string
}

// startCommand 开机命令
type startCommand struct {
	recv receiver
}

func (s *startCommand) excute() string {
	return s.recv.start()
}

func newStartCommand() *startCommand {
	return &startCommand{}
}

// rebootCommand 重启命令
type rebootCommand struct {
	recv receiver
}

func (r *rebootCommand) excute() string {
	return r.recv.reboot()
}

func newRebootCommand() *rebootCommand {
	return &rebootCommand{}
}

// receiver 接收者
type receiver struct{}

func (r *receiver) start() string {
	return "start"
}

func (r *receiver) reboot() string {
	return "reboot"
}

// invoker 调用者
type invoker struct {
	option1 commander
	option2 commander
}

func (i *invoker) excute() string {
	return i.option1.excute() + ":" + i.option2.excute()
}

func newInvoker(op1, op2 commander) *invoker {
	return &invoker{
		option1: op1,
		option2: op2,
	}
}
