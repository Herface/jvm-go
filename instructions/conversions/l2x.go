package conversions

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type L2D struct{ common.NoOperandsInstruction }
type L2F struct{ common.NoOperandsInstruction }
type L2I struct{ common.NoOperandsInstruction }

func (this *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopLong()
	stack.PushDouble(float64(popInt))
}
func (this *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopLong()
	stack.PushFloat(float32(popInt))
}
func (this *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopFloat()
	stack.PushInt(int32(popInt))
}
