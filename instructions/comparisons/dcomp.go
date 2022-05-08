package comparisons

// 浮点数比较

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type DCMPG struct{ common.NoOperandsInstruction }
type DCMPL struct{ common.NoOperandsInstruction }

func _gcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (this *DCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}
func (this *DCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
