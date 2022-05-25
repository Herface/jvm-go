package classfile

// LineNumberTableAttribute 源代码行号表
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (this *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	this.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range this.lineNumberTable {
		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.ReadUint16(),
			lineNumber: reader.ReadUint16(),
		}
	}
}

func (this *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(this.lineNumberTable) - 1; i >= 0; i-- {
		entry := this.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}