package heap

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
