package comparisons

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

// goto 无条件跳转

type GOTO struct{ common.BranchInstruction }

func (this *GOTO) Execute(frame *rtda.Frame) {
	common.Branch(frame, this.Offset)
}
