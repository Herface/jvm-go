package loads

// 加载指令 从本地变量表加载到操作数栈

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD struct{ common.Index8Instruction }
type ILOAD_0 struct{ common.NoOperandsInstruction }
type ILOAD_1 struct{ common.NoOperandsInstruction }
type ILOAD_2 struct{ common.NoOperandsInstruction }
type ILOAD_3 struct{ common.NoOperandsInstruction }

func (this *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, this.Index)
}
func (this *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}
func (this *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (this *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (this *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
