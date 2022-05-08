package math

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
	"math"
)

type DREM struct{ common.NoOperandsInstruction }
type FREM struct{ common.NoOperandsInstruction }
type IREM struct{ common.NoOperandsInstruction }
type LREM struct{ common.NoOperandsInstruction }

func (this *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}
func (this *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

func (this *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := math.Mod(float64(v1), float64(v2))
	stack.PushFloat(float32(result))
}

func (this *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
