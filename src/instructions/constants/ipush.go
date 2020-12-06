package constants

import "gvm/src/instructions/base"
import "gvm/src/rtda"

type BIPUSH struct{ val int8 }
type SIPUSH struct{ val int16 }

func (self *BIPUSH) FetchOperand(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Excute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperand(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Excute(frame *rtda.Frame) {
	i := int64(self.val)
	frame.OperandStack().PushLong(i)
}
