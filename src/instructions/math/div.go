package math

import (
	"gvm/src/instructions/base"
	"gvm/src/rtda"
)

type DDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v1 / v2)
}

type FDIV struct{ base.NoOperandsInstruction }

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v1 / v2)
}

type LDIV struct{ base.NoOperandsInstruction }

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushLong(v1 / v2)
}

type IDIV struct{ base.NoOperandsInstruction }

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushInt(v1 / v2)
}
