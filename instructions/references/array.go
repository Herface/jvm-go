package references
// 数组相关指令

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

// 数组类型
const (
	AT_BOOLEAN = 4
	AT_CHAR = 5
	AT_FLOAT = 6
	AT_DOUBLE = 7
	AT_BYTE = 8
	AT_SHORT = 9
	AT_INT = 10
	AT_LONG = 11
)
// NEW_ARRAY 基本类型数组
type NEW_ARRAY struct {
	atype uint8
}
// ANEW_ARRAY 引用类型数组
type ANEW_ARRAY struct{ common.Index16Instruction }
// ARRAY_LENGTH 数组长度
type ARRAY_LENGTH struct{ common.NoOperandsInstruction }


func (this *NEW_ARRAY) FetchOperands(reader *common.BytecodeReader) {
	this.atype = reader.ReadUint8()
}
func (this *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func (this *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}

func (this *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, this.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}