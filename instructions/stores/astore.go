package stores

// 存储指令 弹出操作数栈顶 存入本地变量表

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

type ASTORE struct{ common.Index8Instruction }
type ASTORE_0 struct{ common.NoOperandsInstruction }
type ASTORE_1 struct{ common.NoOperandsInstruction }
type ASTORE_2 struct{ common.NoOperandsInstruction }
type ASTORE_3 struct{ common.NoOperandsInstruction }

func (this *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, this.Index)
}
func (this *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}
func (this *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (this *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (this *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
