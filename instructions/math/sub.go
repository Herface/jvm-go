package math

// 减法指令
import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type ISUB struct{ common.NoOperandsInstruction }
type LSUB struct{ common.NoOperandsInstruction }
type FSUB struct{ common.NoOperandsInstruction }
type DSUB struct{ common.NoOperandsInstruction }

func (this *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

func (this *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	result := v1 - v2
	stack.PushLong(result)
}
func (this *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	result := v1 - v2
	stack.PushFloat(result)
}

func (this *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	result := v1 - v2
	stack.PushDouble(result)
}
