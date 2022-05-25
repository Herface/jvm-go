package loads

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

type AALOAD struct{ common.NoOperandsInstruction }
type BALOAD struct{ common.NoOperandsInstruction }
type CALOAD struct{ common.NoOperandsInstruction }
type DALOAD struct{ common.NoOperandsInstruction }
type FALOAD struct{ common.NoOperandsInstruction }
type IALOAD struct{ common.NoOperandsInstruction }
type LALOAD struct{ common.NoOperandsInstruction }
type SALOAD struct{ common.NoOperandsInstruction }


func (this *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

func (this *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Bytes()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}

func (this *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Bytes()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}

func (this *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Doubles()
	checkIndex(len(refs), index)
	stack.PushDouble(refs[index])
}

func (this *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Floats()
	checkIndex(len(refs), index)
	stack.PushFloat(refs[index])
}

func (this *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Ints()
	checkIndex(len(refs), index)
	stack.PushInt(refs[index])
}

func (this *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Longs()
	checkIndex(len(refs), index)
	stack.PushLong(refs[index])
}

func (this *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Shorts()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}

func checkIndex(length int, index int32) {
	if int(index) >= length {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}