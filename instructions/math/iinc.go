package math

// 自增指令
import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type IINC struct {
	Index uint
	Const int
}

func (this *IINC) FetchOperands(reader *common.BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
	this.Const = int(reader.ReadInt8())
}
func (this *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	i := localVars.GetInt(this.Index)
	localVars.SetInt(this.Index, i+int32(this.Const))
}
