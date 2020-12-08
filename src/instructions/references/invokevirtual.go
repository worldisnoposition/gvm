package references

import (
	"fmt"
	"gvm/src/instructions/base"
	"gvm/src/rtda"
	"gvm/src/rtda/heap"
)

type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(B)V":
			fmt.Printf("%c\n", stack.PopInt() != 0)
		case "(C)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(S)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(I)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(J)V":
			fmt.Printf("%v\n", stack.PopLong() != 0)
		case "(F)V":
			fmt.Printf("%v\n", stack.PopFloat() != 0)
		case "(D)V":
			fmt.Printf("%v\n", stack.PopDouble() != 0)

		default:
			panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}
