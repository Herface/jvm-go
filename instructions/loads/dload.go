package loads

// double load 指令
import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

type DLOAD struct{ common.Index8Instruction }
type DLOAD_0 struct{ common.NoOperandsInstruction }
type DLOAD_1 struct{ common.NoOperandsInstruction }
type DLOAD_2 struct{ common.NoOperandsInstruction }
type DLOAD_3 struct{ common.NoOperandsInstruction }

func (this DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, this.Index)
}

func (this DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (this *DLOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (this *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (this *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
