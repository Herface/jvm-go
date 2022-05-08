package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (this *ClassReader) ReadUint8() uint8 {
	b := this.data[0]
	this.data = this.data[1:]
	return b
}

func (this *ClassReader) ReadUint16() uint16 {
	u := binary.BigEndian.Uint16(this.data[:2])
	this.data = this.data[2:]
	return u
}

func (this *ClassReader) ReadUint16s() []uint16 {
	length := this.ReadUint16()
	data := make([]uint16, length)
	for i := range data {
		data[i] = this.ReadUint16()
	}
	return data
}
func (this *ClassReader) ReadUint32() uint32 {
	bytes := this.data[:4]
	this.data = this.data[4:]
	return binary.BigEndian.Uint32(bytes)
}
func (this *ClassReader) ReadUint64() uint64 {
	bytes := this.data[:8]
	this.data = this.data[8:]
	return binary.BigEndian.Uint64(bytes)
}
func (this *ClassReader) ReadBytes(n int) []byte {
	bytes := this.data[:n]
	this.data = this.data[n:]
	return bytes
}

func NewClassReader(data []byte) *ClassReader {
	return &ClassReader{
		data: data,
	}
}
