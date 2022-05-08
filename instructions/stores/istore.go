package stores

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

type ISTORE struct{ common.Index8Instruction }
type ISTORE_0 struct{ common.NoOperandsInstruction }
type ISTORE_1 struct{ common.NoOperandsInstruction }
type ISTORE_2 struct{ common.NoOperandsInstruction }
type ISTORE_3 struct{ common.NoOperandsInstruction }

func (this *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, this.Index)
}
func (this *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}
func (this *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (this *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (this *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
