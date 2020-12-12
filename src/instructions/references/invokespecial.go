package references

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
	"gvm/src/rtda/heap"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolveClass := methodRef.ResolveClass()
	resolveMethod := methodRef.ResolveMethod()
	if resolveMethod.Name() == "<init>" && resolveMethod.Class() != resolveClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolveMethod.IsStatic() {
		panic("java.lang.InCompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(resolveMethod.ArgSlotCount())
	if ref != nil {
		panic("java.lang.NullPointException")
	}
	if resolveMethod.IsProtected() && resolveMethod.Class().IsSuperClassOf(currentClass) && resolveMethod.Class().GetPackageName() != currentClass.GetPackageName() && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := resolveMethod
	if currentClass.IsSuper() && resolveClass.IsSuperClassOf(currentClass) && resolveMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lnag.AbstractMethodError")
	}
}
