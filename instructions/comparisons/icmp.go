package comparisons

// 比较指令
// 栈顶两个值比较 结果入栈
import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)

type LCMP struct{ common.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
