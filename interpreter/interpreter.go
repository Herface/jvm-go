package interpreter

import (
	"jvm-go/classfile"
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func Interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxStack := codeAttr.MaxStack()
	maxLocals := codeAttr.MaxLocals()
	code := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)
	exec(thread, code)

}

func exec(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := common.NewByteCodeReader(bytecode)
	for {
		pc := frame.GetNextPC()
		thread.SetPc(pc)
		reader.Reset(bytecode, pc)
		op := common.NewInstruction(reader.ReadUint8())
		op.FetchOperands(reader)
		// 预计的下一条指令 可能会在执行过程中改变
		frame.SetNextPC(reader.PC())
		op.Execute(frame)
	}
}
