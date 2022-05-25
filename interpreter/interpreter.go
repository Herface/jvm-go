package interpreter

import (
	"jvm-go/instructions/common"
	"jvm-go/instructions/factory"
	"jvm-go/rtda"
)

func Interpret(thread *rtda.Thread) {
	exec(thread)
}

func exec(thread *rtda.Thread) {
	currentFrame := thread.CurrentFrame()
	method := currentFrame.Method()
	reader := common.NewByteCodeReader(method.Code())
	for {
		currentFrame = thread.CurrentFrame()
		thread.SetPc(currentFrame.GetNextPC())
		reader.Reset(currentFrame.Method().Code(), thread.PC())
		op := factory.NewInstruction(reader.ReadUint8())
		op.FetchOperands(reader)
		// 预计的下一条指令 可能会在执行过程中改变
		currentFrame.SetNextPC(reader.PC())
		op.Execute(currentFrame)
		if thread.IsEmpty() {
			break
		}
	}
}
