package classfile

// ConstantPool 常量池
type ConstantPool []ConstantInfo

func (this ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := this[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func readConstantPool(reader *ClassReader) ConstantPool {
	n := int(reader.ReadUint16())
	cp := make([]ConstantInfo, n)
	for i := 1; i < n; i++ {
		info := ReadConstantInfo(reader, cp)
		cp[i] = info
		switch info.(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}
func (this ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := this.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := this.getUtf8(ntInfo.nameIndex)
	aType := this.getUtf8(ntInfo.descriptorIndex)
	return name, aType
}

func (this ConstantPool) getClassName(index uint16) string {
	classInfo := this.getConstantInfo(index).(*ConstantClassInfo)
	return this.getUtf8(classInfo.nameIndex)
}
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.val
}
