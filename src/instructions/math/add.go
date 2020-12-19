package math

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type IADD struct{ base.NoOperandsInstruction }
type DADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 + v1
	stack.PushInt(result)
}
func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v2 + v1
	stack.PushFloat(result)
}
func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v2 + v1
	stack.PushDouble(result)
}
func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 + v1
	stack.PushLong(result)
}
