package heap

import (
	"gvm/src/classfile"
	"jvmgo/ch06/rtda/heap"
)

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Name() string {
	return self.name
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) isAccessibleTo(class *Class) bool {
	return false //todo
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) calcArgSlotCount() {
	parseedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parseedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&heap.ACC_ABSTRACT

}

func newMethods(class *Class, cfMthods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMthods))
	for i, cfMethod := range cfMthods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}
