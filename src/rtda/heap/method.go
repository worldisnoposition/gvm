package heap

import "gvm/src/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
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
