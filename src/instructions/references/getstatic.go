package references

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	//stack := frame.OperandStack()
	//stack.
}
