package comparisons

// 浮点数比较

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type FCMPG struct{ common.NoOperandsInstruction }
type FCMPL struct{ common.NoOperandsInstruction }

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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

func (this *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}
func (this *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
