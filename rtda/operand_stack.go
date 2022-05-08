package rtda

import (
	"jvm-go/rtda/heap"
	"math"
)

// OperandStack 操作数栈
type OperandStack struct {
	size  uint
	slots []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}
func (this *OperandStack) PushInt(val int32) {
	this.slots[this.size].num = val
	this.size++
}
func (this *OperandStack) PopInt() int32 {
	this.size--
	return this.slots[this.size].num
}

func (this *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	this.slots[this.size].num = int32(bits)
	this.size++
}
func (this *OperandStack) PopFloat() float32 {
	this.size--
	return math.Float32frombits(uint32(this.slots[this.size].num))
}

func (this *OperandStack) PushLong(val int64) {
	this.slots[this.size].num = int32(val)
	this.slots[this.size+1].num = int32(val >> 32)
	this.size += 2
}

func (this *OperandStack) PopLong() int64 {
	this.size -= 2
	low := uint32(this.slots[this.size].num)
	high := uint32(this.slots[this.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (this *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	this.PushLong(int64(bits))
}
func (this *OperandStack) PopDouble() float64 {
	bits := uint64(this.PopLong())
	return math.Float64frombits(bits)
}

func (this *OperandStack) PushRef(ref *heap.Object) {
	this.slots[this.size].ref = ref
	this.size++
}
func (this *OperandStack) PopRef() *heap.Object {
	this.size--
	ref := this.slots[this.size].ref
	this.slots[this.size].ref = nil // help gc
	return ref
}
func (this *OperandStack) PushSlot(slot Slot) {
	this.slots[this.size] = slot
	this.size++
}
func (this *OperandStack) PopSlot() Slot {
	this.size--
	return this.slots[this.size]
}

func (this *OperandStack) PeekSlot() Slot {
	return this.slots[this.size-1]
}
func (this *OperandStack) InsertDownN(slot Slot, downN uint) {
	i := this.size - 2
	index := this.size - 1 - downN
	for ; i >= index; i-- {
		this.slots[i+1] = this.slots[i]
	}
	i++
	this.slots[i] = slot
}

func (this *OperandStack) Get(index uint) Slot {
	return this.slots[index]
}
