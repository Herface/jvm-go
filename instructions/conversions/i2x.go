package conversions

// int 转其他类型

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type I2B struct{ common.NoOperandsInstruction }
type I2C struct{ common.NoOperandsInstruction }
type I2D struct{ common.NoOperandsInstruction }
type I2F struct{ common.NoOperandsInstruction }
type I2L struct{ common.NoOperandsInstruction }
type I2S struct{ common.NoOperandsInstruction }

func (this *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopInt()
	stack.PushInt(int32(int8(popInt)))
}
func (this *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopInt()
	stack.PushInt(int32(uint16(popInt)))
}
func (this *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopInt()
	stack.PushDouble(float64(popInt))
}

func (this *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopInt()
	stack.PushFloat(float32(popInt))
}

func (this *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopInt()
	stack.PushLong(int64(popInt))
}

func (this *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	popInt := stack.PopInt()
	stack.PushInt(int32(int16(popInt)))
}
