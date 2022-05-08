package comparisons

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IFEQ struct{ common.BranchInstruction }
type IFNE struct{ common.BranchInstruction }
type IFLT struct{ common.BranchInstruction }
type IFLE struct{ common.BranchInstruction }
type IFGT struct{ common.BranchInstruction }
type IFGE struct{ common.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		common.Branch(frame, self.Offset)
	}
}
func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		common.Branch(frame, self.Offset)
	}
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		common.Branch(frame, self.Offset)
	}
}
func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		common.Branch(frame, self.Offset)
	}
}
func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		common.Branch(frame, self.Offset)
	}
}
func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		common.Branch(frame, self.Offset)
	}
}
