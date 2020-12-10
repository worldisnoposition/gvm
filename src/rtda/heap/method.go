package heap

import "gvm/src/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
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

func newMethods(class *Class, cfMthods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMthods))
	for i, cfMethod := range cfMthods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
	}
	return methods
}
