package math

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IADD struct{ common.NoOperandsInstruction }
type LADD struct{ common.NoOperandsInstruction }
type FADD struct{ common.NoOperandsInstruction }
type DADD struct{ common.NoOperandsInstruction }

func (this *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (this *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	result := v1 + v2
	stack.PushLong(result)
}
func (this *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	result := v1 + v2
	stack.PushFloat(result)
}

func (this *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	result := v1 + v2
	stack.PushDouble(result)
}
