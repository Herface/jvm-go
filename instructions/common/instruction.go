package common

import "jvm-go/rtda"

type Instruction interface {
	// FetchOperands 获取操作码
	FetchOperands(reader *BytecodeReader)
	// Execute 执行指令
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction 0x0 无操作指令
type NoOperandsInstruction struct{}

// FetchOperands 读取操作码
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {}

func (this *NoOperandsInstruction) Execute(frame *rtda.Frame) {}

// BranchInstruction 分支指令 跳转到指定pc + offset
type BranchInstruction struct {
	Offset int
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset = int(reader.ReadInt16())
}

// Index8Instruction  存取本地变量表指令 1字节操作码
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction 读常量池指令 2字节操作码
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
