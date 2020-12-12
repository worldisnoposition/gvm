package references

import (
	"fmt"
	"gvm/src/instructions/base"
	"gvm/src/rtda"
	"gvm/src/rtda/heap"
)

type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolveMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointEXception")
	}
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) && resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() && ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(C)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())

	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}

//func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
//	cp := frame.Method().Class().ConstantPool()
//	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
//	if methodRef.Name() == "println" {
//		stack := frame.OperandStack()
//		switch methodRef.Descriptor() {
//		case "(Z)V":
//			fmt.Printf("%v\n", stack.PopInt())
//		case "(B)V":
//			fmt.Printf("%c\n", stack.PopInt())
//		case "(C)V":
//			fmt.Printf("%v\n", stack.PopInt())
//		case "(S)V":
//			fmt.Printf("%v\n", stack.PopInt())
//		case "(I)V":
//			fmt.Printf("%v\n", stack.PopInt())
//		case "(J)V":
//			fmt.Printf("%v\n", stack.PopLong())
//		case "(F)V":
//			fmt.Printf("%v\n", stack.PopFloat())
//		case "(D)V":
//			fmt.Printf("%v\n", stack.PopDouble())
//
//		default:
//			panic("println: " + methodRef.Descriptor())
//		}
//		stack.PopRef()
//	}
//}
