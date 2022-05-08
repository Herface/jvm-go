package stores

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

type DSTORE struct{ common.Index8Instruction }
type DSTORE_0 struct{ common.NoOperandsInstruction }
type DSTORE_1 struct{ common.NoOperandsInstruction }
type DSTORE_2 struct{ common.NoOperandsInstruction }
type DSTORE_3 struct{ common.NoOperandsInstruction }

func (this *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, this.Index)
}
func (this *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}
func (this *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

func (this *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

func (this *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
