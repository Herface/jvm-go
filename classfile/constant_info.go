package classfile

const (
	ConstantClass              = 7
	ConstantFieldref           = 9
	ConstantMethodref          = 10
	ConstantInterfacemethodref = 11
	ConstantString             = 8
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantNameandtype        = 12
	ConstantUtf8               = 1
	ConstantMethodhandle       = 15
	ConstantMethodtype         = 16
	ConstantInvokedynamic      = 18
)

// ConstantInfo 常量池项
type ConstantInfo interface {
	ReadInfo(reader *ClassReader)
}

func ReadConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.ReadUint8()
	c := newConstantInfo(tag, cp)
	c.ReadInfo(reader)
	return c

}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case ConstantInteger:
		return &ConstantIntegerInfo{}
	case ConstantFloat:
		return &ConstantFloatInfo{}
	case ConstantLong:
		return &ConstantLongInfo{}
	case ConstantDouble:
		return &ConstantDoubleInfo{}
	case ConstantUtf8:
		return &ConstantUtf8Info{}
	case ConstantString:
		return &ConstantStringInfo{cp: cp}
	case ConstantClass:
		return &ConstantClassInfo{cp: cp}
	case ConstantFieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantMethodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantInterfacemethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantNameandtype:
		return &ConstantNameAndTypeInfo{}
	case ConstantMethodtype:
		return &ConstantMethodTypeInfo{}
	case ConstantMethodhandle:
		return &ConstantMethodHandleInfo{}
	case ConstantInvokedynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
