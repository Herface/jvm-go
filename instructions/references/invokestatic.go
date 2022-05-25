package references

// 静态方法调用指令

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
)

type INVOKE_STATIC struct{ common.Index16Instruction }

func (this *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(this.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		// 执行<clinit>方法 保存当前帧pc
		frame.RevertNextPC()
		common.InitClass(frame.Thread(), class)
		return
	}
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	common.InvokeMethod(frame, resolvedMethod)
}
