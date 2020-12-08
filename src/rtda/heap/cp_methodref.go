package heap

import "gvm/src/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func (self MethodRef) Name() string {
	return self.name
}

func (self MethodRef) Descriptor() string {
	return self.descriptor
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
