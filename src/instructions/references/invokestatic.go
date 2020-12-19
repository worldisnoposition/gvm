package references

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
	"gvm/src/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolveMethod := methodRef.ResolveMethod()
	if !resolveMethod.IsStatic() {
		panic("java.lang.IncompatibleClassError")
	}
	class := resolveMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame, resolveMethod)
}
