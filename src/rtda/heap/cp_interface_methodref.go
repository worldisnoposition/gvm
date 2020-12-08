package heap

import "gvm/src/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMehtodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
