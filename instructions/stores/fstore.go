package stores

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

type FSTORE struct{ common.Index8Instruction }
type FSTORE_0 struct{ common.NoOperandsInstruction }
type FSTORE_1 struct{ common.NoOperandsInstruction }
type FSTORE_2 struct{ common.NoOperandsInstruction }
type FSTORE_3 struct{ common.NoOperandsInstruction }

func (this *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, this.Index)
}
func (this *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}
func (this *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

func (this *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

func (this *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
