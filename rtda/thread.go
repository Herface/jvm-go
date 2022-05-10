package rtda

import "jvm-go/rtda/heap"

type Thread struct {
	pc     int
	stack  *Stack
	method *heap.Method
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

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}

func (this *Thread) IsEmpty() bool {
	return this.stack.IsEmpty()
}
