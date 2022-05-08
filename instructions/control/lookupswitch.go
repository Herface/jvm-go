package control

// lookupswitch 查找匹配

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (this *LOOKUP_SWITCH) FetchOperands(reader *common.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.npairs = reader.ReadInt32()
	this.matchOffsets = reader.ReadInt32s(this.npairs * 2)
}

func (this *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < this.npairs*2; i += 2 {
		if this.matchOffsets[i] == key {
			offset := this.matchOffsets[i+1]
			common.Branch(frame, int(offset))
			return
		}
	}
	common.Branch(frame, int(this.defaultOffset))
}
