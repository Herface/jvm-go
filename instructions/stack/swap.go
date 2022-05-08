package stack

// 交换指令 交换栈顶两个元素

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type SWAP struct{ common.NoOperandsInstruction }

func (this *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
