package comparisons

// 比较跳转指令

// 比较栈顶两值 执行跳转

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IF_ICMPEQ struct{ common.BranchInstruction }
type IF_ICMPNE struct{ common.BranchInstruction }
type IF_ICMPLT struct{ common.BranchInstruction }
type IF_ICMPLE struct{ common.BranchInstruction }
type IF_ICMPGT struct{ common.BranchInstruction }
type IF_ICMPGE struct{ common.BranchInstruction }

func (this *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		common.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		common.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		common.Branch(frame, this.Offset)
	}
}

func (this *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		common.Branch(frame, this.Offset)
	}
}
func (this *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		common.Branch(frame, this.Offset)
	}
}
func (this *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		common.Branch(frame, this.Offset)
	}
}
