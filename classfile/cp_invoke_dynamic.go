package classfile

// 动态调用常量信息

type ConstantMethodHandleInfo struct {
	referenceKind  byte
	referenceIndex uint16
}

func (this *ConstantMethodHandleInfo) ReadInfo(reader *ClassReader) {
	this.referenceKind = reader.ReadUint8()
	this.referenceIndex = reader.ReadUint16()
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (this *ConstantMethodTypeInfo) ReadInfo(reader *ClassReader) {
	this.descriptorIndex = reader.ReadUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16

	nameAndTypeIndex uint16
}

func (this *ConstantInvokeDynamicInfo) ReadInfo(reader *ClassReader) {
	this.bootstrapMethodAttrIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
}
