package heap

import "sync"

type Object struct {
	class *Class
	//fields Slots
	data interface{}
	// extra magic
	extra interface{}
	// 支持synchronized
	mutex sync.Mutex
	cond  sync.Cond
}

func (this *Object) Fields() Slots {
	return this.data.(Slots)
}
func (this *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(this.class)
}
func (this *Object) Class() *Class {
	return this.class
}

func (this *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := this.class.getField(name, descriptor, false)
	slots := this.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (this *Object) Extra() interface{} {
	return this.extra
}

func (this *Object) GetRefVar(name, descriptor string) *Object {
	field := this.class.getField(name, descriptor, false)
	slots := this.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (this *Object) SetExtra(extra interface{}) {
	this.extra = extra
}

func (this *Object) Clone() *Object {
	return &Object{
		class: this.class,
		data:  this.cloneData(),
	}
}
func (this *Object) cloneData() interface{} {
	switch this.data.(type) {
	case []int8:
		arr := this.data.([]int8)
		newArr := make([]int8, len(arr))
		copy(newArr, arr)
		return newArr
	case []int16:
		arr := this.data.([]int16)
		newArr := make([]int16, len(arr))
		copy(newArr, arr)
		return newArr
	case []uint16:
		arr := this.data.([]uint16)
		newArr := make([]uint16, len(arr))
		copy(newArr, arr)
		return newArr
	case []int32:
		arr := this.data.([]int32)
		newArr := make([]int32, len(arr))
		copy(newArr, arr)
		return newArr
	case []int64:
		arr := this.data.([]int64)
		newArr := make([]int64, len(arr))
		copy(newArr, arr)
		return newArr
	case []float32:
		arr := this.data.([]float32)
		newArr := make([]float32, len(arr))
		copy(newArr, arr)
		return newArr
	case []float64:
		arr := this.data.([]float64)
		newArr := make([]float64, len(arr))
		copy(newArr, arr)
		return newArr
	case []*Object:
		elements := this.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default: // []Slot
		slots := this.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}
