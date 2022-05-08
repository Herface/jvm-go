package loads

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

type LLOAD struct{ common.Index8Instruction }
type LLOAD_0 struct{ common.NoOperandsInstruction }
type LLOAD_1 struct{ common.NoOperandsInstruction }
type LLOAD_2 struct{ common.NoOperandsInstruction }
type LLOAD_3 struct{ common.NoOperandsInstruction }

func (this LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, this.Index)
}

func (this LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (this *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (this *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (this *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
