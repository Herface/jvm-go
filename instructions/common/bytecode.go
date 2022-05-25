package common

// BytecodeReader 读取字节码指令
type BytecodeReader struct {
	code []byte
	pc   int
}

func (this *BytecodeReader) ReadUint8() uint8 {
	i := this.code[this.pc]
	this.pc++
	return i
}

func (this *BytecodeReader) Reset(code []byte, pc int) {
	this.code = code
	this.pc = pc
}

func (this *BytecodeReader) ReadInt8() int8 {
	return int8(this.ReadUint8())
}

func (this *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(this.ReadUint8())
	byte2 := uint16(this.ReadUint8())
	return (byte1 << 8) | byte2
}

func (this *BytecodeReader) ReadInt16() int16 {
	return int16(this.ReadUint16())
}
func (this *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(this.ReadUint8())
	byte2 := int32(this.ReadUint8())
	byte3 := int32(this.ReadUint8())
	byte4 := int32(this.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (this *BytecodeReader) SkipPadding() {
	for this.pc%4 != 0 {
		this.ReadUint8()
	}
}
func (this *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = this.ReadInt32()
	}
	return ints
}

func (this *BytecodeReader) PC() int {
	return this.pc
}

func NewByteCodeReader(code []byte) *BytecodeReader {
	return &BytecodeReader{code: code}
}
