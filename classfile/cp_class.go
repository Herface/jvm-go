package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (this *ConstantClassInfo) ReadInfo(reader *ClassReader) {
	this.nameIndex = reader.ReadUint16()
}
func (this *ConstantClassInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}
