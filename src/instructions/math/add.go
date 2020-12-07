package math

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type IADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 + v1
	stack.PushInt(result)
}
