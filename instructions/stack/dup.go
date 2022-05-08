package stack

// 栈顶复制指令

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type DUP struct{ common.NoOperandsInstruction }
type DUP_X1 struct{ common.NoOperandsInstruction }
type DUP_X2 struct{ common.NoOperandsInstruction }
type DUP2 struct{ common.NoOperandsInstruction }
type DUP2_X1 struct{ common.NoOperandsInstruction }
type DUP2_X2 struct{ common.NoOperandsInstruction }

func insertToN() {

}

func (this *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

func (this *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PeekSlot()
	stack.PushSlot(slot)
	stack.InsertDownN(slot, 2)
}

func (this *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PeekSlot()
	stack.PushSlot(slot)
	stack.InsertDownN(slot, 3)
}

func (this *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PeekSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

func (this *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot2)
	stack.InsertDownN(slot2, 3)
	stack.PushSlot(slot1)
	stack.InsertDownN(slot2, 3)

}

func (this *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot2)
	stack.InsertDownN(slot2, 4)
	stack.PushSlot(slot1)
	stack.InsertDownN(slot2, 4)

}
