package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (this *UnparsedAttribute) readInfo(reader *ClassReader) {
	this.info = reader.ReadBytes(int(this.length))
}
