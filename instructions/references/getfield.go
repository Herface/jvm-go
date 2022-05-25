package references
// 获取对象属性指令

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

type GUT_FIELD struct{ common.Index16Instruction }

func (self *GUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		stack.PushInt(ref.Fields().GetInt(slotId))
	case 'F':
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		stack.PushFloat(ref.Fields().GetFloat(slotId))
	case 'J':
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		stack.PushLong(ref.Fields().GetLong(slotId))
	case 'D':
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		stack.PushDouble(ref.Fields().GetDouble(slotId))
	case 'L', '[':
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		stack.PushRef(ref.Fields().GetRef(slotId))
	}
}
