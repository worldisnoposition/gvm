package classfile

import (
	"gvm/src/rtda/heap"
	"math"
)

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value() heap.Constant {
	return self.val
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantFloatInfo) Value() heap.Constant {
	return self.val
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantLongInfo) Value() heap.Constant {
	return self.val
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() heap.Constant {
	return self.val
}
