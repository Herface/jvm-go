package reversed


import "jvm-go/native"
import _ "jvm-go/native/java/lang"

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
)
// INVOKE_NATIVE 调用本地方法指令
type INVOKE_NATIVE struct{ common.NoOperandsInstruction }

func (this *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}

