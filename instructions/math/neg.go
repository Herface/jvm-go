package math

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type INEG struct{ common.NoOperandsInstruction }
type LNEG struct{ common.NoOperandsInstruction }
type FNEG struct{ common.NoOperandsInstruction }
type DNEG struct{ common.NoOperandsInstruction }

func (this *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	stack.PushInt(-v)
}

func (this *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopLong()
	stack.PushLong(-v)
}

func (this *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopFloat()
	stack.PushFloat(-v)
}

func (this *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopDouble()
	stack.PushDouble(-v)
}
