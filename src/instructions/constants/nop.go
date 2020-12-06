package constants

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {

}
