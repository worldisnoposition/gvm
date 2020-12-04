package classfile

type ConstanClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstanClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstanClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
