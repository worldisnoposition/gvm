package base

import "gvm/src/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(Frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

type BranchInstruction struct {
	Offset int
}

type Index8Instruction struct {
	Index uint
}

type Index16Instruction struct {
	Index uint
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt8())
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt16())
}
