package loads

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.Index8Instruction }
type ALOAD_1 struct{ base.Index8Instruction }
type ALOAD_2 struct{ base.Index8Instruction }
type ALOAD_3 struct{ base.Index8Instruction }

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}
func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}
func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}
func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}
func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
