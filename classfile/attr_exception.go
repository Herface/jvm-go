package classfile

// ExceptionsAttribute 方法声明异常表
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (this *ExceptionsAttribute) readInfo(reader *ClassReader) {
	this.exceptionIndexTable = reader.ReadUint16s()
}
func (this *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return this.exceptionIndexTable
}
