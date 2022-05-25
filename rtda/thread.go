package rtda

import "jvm-go/rtda/heap"

type Thread struct {
	pc     int
	stack  *Stack
	method *heap.Method
	jlthread *heap.Object
}

func NewThread() *Thread {
	return &Thread{
		stack: NewStack(1024),
	}
}

func (this *Thread) PC() int {
	return this.pc
}
func (this *Thread) SetPc(pc int) {
	this.pc = pc
}
func (this *Thread) PushFrame(frame *Frame) {
	this.stack.push(frame)
}
func (this *Thread) PopFrame() *Frame {
	return this.stack.pop()
}
func (this *Thread) CurrentFrame() *Frame {
	return this.stack.peek()
}

func (this *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(this, method)
}

func (this *Thread) IsEmpty() bool {
	return this.stack.IsEmpty()
}
func (this *Thread) ClearStack() {
	this.stack.clear()
}

func (this *Thread) IsStackEmpty() bool {
	return this.IsEmpty()
}

func (this *Thread) GetFrames() []*Frame {
	return this.stack.getFrames()
}

func (this *Thread) JLThread() *heap.Object {
	return this.jlthread
}

func (this *Thread) SetJLThread(jlthread *heap.Object) {
	this.jlthread = jlthread
}