package stores

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

type LSTORE struct{ common.Index8Instruction }
type LSTORE_0 struct{ common.NoOperandsInstruction }
type LSTORE_1 struct{ common.NoOperandsInstruction }
type LSTORE_2 struct{ common.NoOperandsInstruction }
type LSTORE_3 struct{ common.NoOperandsInstruction }

func (this *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, this.Index)
}
func (this *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}
func (this *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (this *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (this *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
