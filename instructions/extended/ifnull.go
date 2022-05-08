package extended

// 判空指令

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IFNULL struct{ common.BranchInstruction }
type IFNONNULL struct{ common.BranchInstruction }

func (this *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		common.Branch(frame, this.Offset)
	}
}

func (this *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		common.Branch(frame, this.Offset)
	}
}
