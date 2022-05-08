package math

// 除法指令
import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IDIV struct{ common.NoOperandsInstruction }
type LDIV struct{ common.NoOperandsInstruction }
type FDIV struct{ common.NoOperandsInstruction }
type DDIV struct{ common.NoOperandsInstruction }

func (this *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (this *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	result := v1 + v2
	stack.PushLong(result)
}
func (this *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	result := v1 + v2
	stack.PushFloat(result)
}

func (this *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	result := v1 + v2
	stack.PushDouble(result)
}
