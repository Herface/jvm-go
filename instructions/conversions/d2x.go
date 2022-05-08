package conversions

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type D2F struct{ common.NoOperandsInstruction }
type D2I struct{ common.NoOperandsInstruction }
type D2L struct{ common.NoOperandsInstruction }

func (this *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopDouble()
	stack.PushFloat(float32(popInt))
}
func (this *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopDouble()
	stack.PushInt(int32(popInt))
}
func (this *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopDouble()
	stack.PushLong(int64(popInt))
}
