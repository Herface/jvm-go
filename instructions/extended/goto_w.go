package extended

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type GOTO_W struct {
	offset uint
}

func (this *GOTO_W) FetchOperands(reader *common.BytecodeReader) {
	this.offset = uint(reader.ReadInt32())
}

func (this *GOTO_W) Execute(frame *rtda.Frame) {
	common.Branch(frame, int(this.offset))
}
