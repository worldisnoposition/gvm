package references

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
