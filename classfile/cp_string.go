package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (this *ConstantStringInfo) ReadInfo(reader *ClassReader) {
	this.stringIndex = reader.ReadUint16()
}

func (this *ConstantStringInfo) String() string {
	return this.cp.getUtf8(this.stringIndex)
}
