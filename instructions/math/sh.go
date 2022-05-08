package math

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type ISHL struct{ common.NoOperandsInstruction }
type ISHR struct{ common.NoOperandsInstruction }
type IUSHR struct{ common.NoOperandsInstruction }
type LSHL struct{ common.NoOperandsInstruction }
type LSHR struct{ common.NoOperandsInstruction }
type LUSHR struct{ common.NoOperandsInstruction }

func (this *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// 32bit 最大位移
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}
func (this *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// 32bit 最大位移
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}
func (this *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (this *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

func (this *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

func (this *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
