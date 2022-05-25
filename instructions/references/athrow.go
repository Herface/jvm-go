package references

import (
	"jvm-go/instructions/common"
	"jvm-go/rtda"
	"jvm-go/rtda/heap"
	"reflect"
)

// ATHROW 异常抛出指令
type ATHROW struct{ common.NoOperandsInstruction }

func (this *ATHROW) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}
	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}
func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		// 当前pc
		pc := frame.GetNextPC() - 1
		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPC > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}
		// 当前栈帧无法处理异常(异常表无法处理异常) 继续找下一个栈帧
		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

// handleUncaughtException 未捕获的异常 输出栈信息
func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()
	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)
	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}
