package classfile

// ConstantValueAttribute 常量属性
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (this *ConstantValueAttribute) readInfo(reader *ClassReader) {
	this.constantValueIndex = reader.ReadUint16()
}
func (this *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return this.constantValueIndex
}
