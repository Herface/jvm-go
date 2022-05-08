package references

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

// 从常量池推入

type LDC struct{ common.Index8Instruction }
type LDC_W struct{ common.Index16Instruction }
type LDC2_W struct{ common.Index16Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)

}
func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}
func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
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
