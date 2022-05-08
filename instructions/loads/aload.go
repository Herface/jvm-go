package loads

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

type ALOAD struct{ common.Index8Instruction }
type ALOAD_0 struct{ common.NoOperandsInstruction }
type ALOAD_1 struct{ common.NoOperandsInstruction }
type ALOAD_2 struct{ common.NoOperandsInstruction }
type ALOAD_3 struct{ common.NoOperandsInstruction }

func (this ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, this.Index)
}

func (this ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (this *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (this *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (this *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
