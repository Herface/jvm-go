package loads

// float load 指令

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

type FLOAD struct{ common.Index8Instruction }
type FLOAD_0 struct{ common.NoOperandsInstruction }
type FLOAD_1 struct{ common.NoOperandsInstruction }
type FLOAD_2 struct{ common.NoOperandsInstruction }
type FLOAD_3 struct{ common.NoOperandsInstruction }

func (this FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, this.Index)
}

func (this FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (this *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (this *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (this *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
