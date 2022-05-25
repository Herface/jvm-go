package lang

import (
	"jvm-go/interpreter"
	"jvm-go/native"
	"jvm-go/rtda"
	"time"
)

func init() {
	native.Register("java/lang/Thread", "start0", "()V", start0)
	native.Register("java/lang/Thread", "currentThread", "()Ljava/lang/Thread;",  currentThread)
	native.Register("java/lang/Thread", "sleep0", "(J)V",  sleep0)
}

func start0(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	class := ref.Class()

	jlThread := class.NewObject()
	this := frame.LocalVars().GetThis()
	method := class.GetMethod("run", "()V")
	thread := rtda.NewThread()
	newFrame := thread.NewFrame(method)
	newFrame.LocalVars().SetRef(0, this)
	thread.PushFrame(newFrame)
	thread.SetJLThread(jlThread)
	go func() {
		interpreter.Interpret(thread)
	}()
}

func currentThread(frame *rtda.Frame) {
	thread := frame.Thread()
	jlThread := thread.JLThread()
	frame.OperandStack().PushRef(jlThread)
}

func sleep0(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	long := localVars.GetLong(1)
	time.Sleep(time.Millisecond * time.Duration(long))
}