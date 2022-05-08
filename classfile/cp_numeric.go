package classfile

import "math"

// 数字类型

type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) ReadInfo(reader *ClassReader) {
	bytes := reader.ReadUint32()
	this.val = int32(bytes)
}

func (this *ConstantIntegerInfo) Value() int32 {
	return this.val
}

type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) ReadInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	this.val = int64(bytes)
}
func (this *ConstantLongInfo) Value() int64 {
	return this.val
}

type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) ReadInfo(reader *ClassReader) {
	bytes := reader.ReadUint32()
	this.val = math.Float32frombits(bytes)
}
func (this *ConstantFloatInfo) Value() float32 {
	return this.val
}

type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) ReadInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	this.val = math.Float64frombits(bytes)
}
func (this *ConstantDoubleInfo) Value() float64 {
	return this.val
}
