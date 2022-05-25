package lang

import (
	"jvm-go/native"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", func(frame *rtda.Frame) {
		this := frame.LocalVars().GetThis()
		interned := heap.InternString(this)
		frame.OperandStack().PushRef(interned)
	})
}

