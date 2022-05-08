package math

// 乘法指令
import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IMUL struct{ common.NoOperandsInstruction }
type LMUL struct{ common.NoOperandsInstruction }
type FMUL struct{ common.NoOperandsInstruction }
type DMUL struct{ common.NoOperandsInstruction }

func (this *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	result := v1 * v2
	stack.PushInt(result)
}

func (this *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	result := v1 * v2
	stack.PushLong(result)
}
func (this *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

func (this *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	result := v1 * v2
	stack.PushDouble(result)
}
