package heap

import (
	"gvm/src/classfile"
	"gvm/src/rtda/heap"
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

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) calcArgSlotCount(paramTypes []string) {
	//parseedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range paramTypes {
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

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} //return
	case 'D':
		self.code = []byte{0xfe, 0xaf} //dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} //freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} //lreturn
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} //areturn
	default:
		self.code = []byte{0xfe, 0xac} //ireturn

	}
}

func newMethods(class *Class, cfMthods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMthods))
	for i, cfMethod := range cfMthods {
		methods[i] = newMethod(class, cfMethod)
		//methods[i] = &Method{}
		//methods[i].class = class
		//methods[i].copyMemberInfo(cfMethod)
		//methods[i].copyAttribute(cfMethod)
		//methods[i].calcArgSlotCount()
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttribute(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}
