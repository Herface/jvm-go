package constants

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

// int 类型以下压入操作数栈 指令

type BIPUSH struct{ val int8 }
type SIPUSH struct{ val int16 }

func (this *BIPUSH) FetchOperands(reader *common.BytecodeReader) {
	this.val = reader.ReadInt8()
}
func (this *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

func (this *SIPUSH) FetchOperands(reader *common.BytecodeReader) {
	this.val = reader.ReadInt16()
}
func (this *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}
