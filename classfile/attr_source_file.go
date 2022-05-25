package classfile
// SourceFileAttribute 源文件名
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (this *SourceFileAttribute) readInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.ReadUint16()
}
func (this *SourceFileAttribute) FileName() string {
	return this.cp.getUtf8(this.sourceFileIndex)
}
