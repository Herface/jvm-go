package classfile

type ConstantUtf8Info struct {
	val string
}

func (self *ConstantUtf8Info) ReadInfo(reader *ClassReader) {
	n := reader.ReadUint16()
	self.val = string(reader.ReadBytes(int(n)))
}
