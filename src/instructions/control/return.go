package control

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	//frame.Thread().PopFrame()
}
