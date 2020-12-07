package math

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type DSUB struct{ base.NoOperandsInstruction }

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v1 - v2)
}

type FSUB struct{ base.NoOperandsInstruction }

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v1 - v2)
}

type LSUB struct{ base.NoOperandsInstruction }

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 - v2)
}

type ISUB struct{ base.NoOperandsInstruction }

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 - v2)
}
