package control

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
