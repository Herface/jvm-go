package comparisons

// 引用类型比较 指令

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IF_ACMPEQ struct{ common.BranchInstruction }
type IF_ACMPNE struct{ common.BranchInstruction }

func (this *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		common.Branch(frame, this.Offset)
	}
}

func (this *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		common.Branch(frame, this.Offset)
	}
}
