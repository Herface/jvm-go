package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}
type Slots []Slot

func newSlots(size uint) Slots {
	if size > 0 {
		return make([]Slot, size)
	}
	return nil
}

func (this Slots) SetInt(index uint, val int32) {
	this[index].num = val
}
func (this Slots) GetInt(index uint) int32 {
	return this[index].num
}

func (this Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	this[index].num = int32(bits)
}
func (this Slots) GetFloat(index uint) float32 {
	bits := uint32(this[index].num)
	return math.Float32frombits(bits)
}

func (this Slots) SetLong(index uint, val int64) {
	this[index].num = int32(val)
	this[index+1].num = int32(val >> 32)
}
func (this Slots) GetLong(index uint) int64 {
	low := uint32(this[index].num)
	high := uint32(this[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (this Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	this.SetLong(index, int64(bits))
}
func (this Slots) GetDouble(index uint) float64 {
	bits := uint64(this.GetLong(index))
	return math.Float64frombits(bits)
}
func (this Slots) SetRef(index uint, ref *Object) {
	this[index].ref = ref
}
func (this Slots) GetRef(index uint) *Object {
	return this[index].ref
}
