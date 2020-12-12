package main

import (
	"fmt"
	"gvm/src/instructions"
	"gvm/src/instructions/base"
	"gvm/src/rtda"
	"gvm/src/rtda/heap"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread, logInst)
	//codeAttr := methodInfo.CodeAttribute()
	//maxLocals := codeAttr.MaxLocals()
	//maxStack := codeAttr.MaxStack()
	//bytecode := codeAttr.Code()
	//
	//thread := rtda.NewThread()
	//frame := thread.NewFrame(maxLocals, maxStack)
	//thread.PushFrame(frame)
	//defer catchErr(frame)
	//loop(thread, bytecode)
}

func loop(thread *rtda.Thread, logInst bool) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	bytecode := frame.Method().Code()
	var pc int
	for {
		pc = frame.NextPC()
		thread.SetPC(pc)
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//fmt.Printf("pc:%2d inst:%T %v", pc, inst, inst)
		//fmt.Printf(" stackSize:%v\n", frame.OperandStack().Size())
		//fmt.Println("	frame.localVars: %v\n", frame.LocalVars())
		if logInst {
			logInstruction(frame, inst)
		}
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
	fmt.Printf(" stackSize:%v", frame.OperandStack().Size())
	fmt.Println("	frame.localVars: %v\n", frame.LocalVars())
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		//fmt.Print("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	if !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
