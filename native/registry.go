package native

import "jvm-go/rtda"

// NativeMethod 本地方法
type NativeMethod func(frame *rtda.Frame)

// registry 本地方法表
var registry = map[string]NativeMethod{}

// Register 注册本地方法
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// 空方法
func emptyNativeMethod(frame *rtda.Frame) {}


// FindNativeMethod 获取本地方法
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}