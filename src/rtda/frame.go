package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func (self Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self Frame) NextPC() int {
	return self.nextPC
}

func (self Frame) Thread() *Thread {
	return self.thread
}

func (self Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

//func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
//	return &Frame{
//		thread:       thread,
//		localVars:    newLocalVars(maxLocals),
//		operandStack: newOperandStack(maxStack),
//	}
//}
