package math

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type DMUL struct{ base.NoOperandsInstruction }

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v2 * v1)
}

type FMUL struct{ base.NoOperandsInstruction }

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v2 * v1)
}

type LMUL struct{ base.NoOperandsInstruction }

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v2 * v1)
}

type IMUL struct{ base.NoOperandsInstruction }

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v2 * v1)
}
