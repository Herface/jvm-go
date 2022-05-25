package rtda

import (
	"jvm-go/rtda/heap"
	"math"
)

// Frame 栈帧 一个方法调用
type Frame struct {
	// 栈顶的下一个栈帧(链表)
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread // 持有帧的线程
	pc           int     // 下一条指令
	method       *heap.Method
}

func (this *Frame) LocalVars() LocalVars {
	return this.localVars
}
func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack
}
func (this *Frame) Thread() *Thread {
	return this.thread
}
func (this *Frame) SetNextPC(pc int) {
	this.pc = pc
}
func (this *Frame) GetNextPC() int {
	return this.pc
}

func (this *Frame) Method() *heap.Method {
	return this.method
}

func (this *Frame) RevertNextPC() {
	this.pc = this.thread.PC()
}


func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    NewLocalVars(method.MaxLocals()),
		operandStack: NewOperandStack(method.MaxStack()),
		method:       method,
	}
}

// Slot 本地变量表 操作数栈 slot
type Slot struct {
	num int32
	ref *heap.Object
}

// LocalVars 本地变量表  一个slot 4字节 double long 占两个slot
type LocalVars []Slot

func NewLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (this LocalVars) SetInt(index uint, val int32) {
	this[index].num = val
}
func (this LocalVars) GetInt(index uint) int32 {
	return this[index].num
}

func (this LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	this[index].num = int32(bits)
}
func (this LocalVars) GetFloat(index uint) float32 {
	bits := uint32(this[index].num)
	return math.Float32frombits(bits)
}

func (this LocalVars) SetLong(index uint, val int64) {
	this[index].num = int32(val)
	this[index+1].num = int32(val >> 32)
}
func (this LocalVars) GetLong(index uint) int64 {
	low := uint32(this[index].num)
	high := uint32(this[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (this LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	this.SetLong(index, int64(bits))
}
func (this LocalVars) GetDouble(index uint) float64 {
	bits := uint64(this.GetLong(index))
	return math.Float64frombits(bits)
}
func (this LocalVars) SetRef(index uint, ref *heap.Object) {
	this[index].ref = ref
}
func (this LocalVars) GetRef(index uint) *heap.Object {
	return this[index].ref
}

func (this LocalVars) SetSlot(index uint, slot Slot) {
	this[index] = slot
}

func (this LocalVars) GetThis() *heap.Object {
	return this.GetRef(0)
}
// Clear 清空操作数栈
func (this *OperandStack) Clear() {
	this.size = 0
	for i := range this.slots {
		this.slots[i].ref = nil
	}
}