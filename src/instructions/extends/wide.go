package extends

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
