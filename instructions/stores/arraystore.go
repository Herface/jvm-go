package stores
// 数组写指令
import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type AASTORE struct{ common.NoOperandsInstruction }
type BASTORE struct{ common.NoOperandsInstruction }
type CASTORE struct{ common.NoOperandsInstruction }
type DASTORE struct{ common.NoOperandsInstruction }
type FASTORE struct{ common.NoOperandsInstruction }
type IASTORE struct{ common.NoOperandsInstruction }
type LASTORE struct{ common.NoOperandsInstruction }
type SASTORE struct{ common.NoOperandsInstruction }



func (this *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Refs()
	common.CheckIndex(len(ints), index)
	ints[index] = val
}

func (this *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Ints()
	common.CheckIndex(len(ints), index)
	ints[index] = val
}

func (this *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Bytes()
	common.CheckIndex(len(ints), index)
	ints[index] = int8(val)
}

func (this *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Chars()
	common.CheckIndex(len(ints), index)
	ints[index] = uint16(val)
}

func (this *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Doubles()
	common.CheckIndex(len(ints), index)
	ints[index] = val
}

func (this *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Floats()
	common.CheckIndex(len(ints), index)
	ints[index] = val
}

func (this *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Longs()
	common.CheckIndex(len(ints), index)
	ints[index] = val
}

func (this *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	common.CheckNotNil(arrRef)
	ints := arrRef.Shorts()
	common.CheckIndex(len(ints), index)
	ints[index] = int16(val)
}