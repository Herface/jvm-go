package classfile

type LocalVariableTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LocalVariableTableAttributeEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (this *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	this.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range this.lineNumberTable {
		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.ReadUint16(),
			lineNumber: reader.ReadUint16(),
		}
	}
}
