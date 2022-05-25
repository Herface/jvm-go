package references

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

// 从常量池推入

type LDC struct{ common.Index8Instruction }
type LDC_W struct{ common.Index16Instruction }
type LDC2_W struct{ common.Index16Instruction }

func (this *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)

}
func (this *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)
}
func (this *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	cp := class.ConstantPool()
	c := cp.GetConstant(this.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	default:
		panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))

	default:
		panic("todo: ldc!")
	}
}
