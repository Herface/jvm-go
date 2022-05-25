package lang

import (
	"jvm-go/native"
	"jvm-go/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Float",
		"floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register("java/lang/Double",
		"doubleToRawIntBits", "(F)I", doubleToRawIntBits)
}
func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}

func doubleToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}