package java

import "gvm/src/native"
import "gvm/src/rtda"

func init() {
	native.Register("java/lang/Obejct", "getClass", "()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
