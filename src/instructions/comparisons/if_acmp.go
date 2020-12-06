package comparisons

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type IF_ACMPEQ struct{ base.BranchInstruction }
type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopRef()
	val1 := stack.PopRef()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopRef()
	val1 := stack.PopRef()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}
