package stack

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type POP struct{ common.NoOperandsInstruction }
type POP2 struct{ common.NoOperandsInstruction }

func (this *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (this *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
