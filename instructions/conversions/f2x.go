package conversions

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type F2D struct{ common.NoOperandsInstruction }
type F2I struct{ common.NoOperandsInstruction }
type F2L struct{ common.NoOperandsInstruction }

func (this *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopFloat()
	stack.PushLong(int64(popInt))
}
func (this *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopFloat()
	stack.PushInt(int32(popInt))
}
func (this *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopFloat()
	stack.PushDouble(float64(popInt))
}
