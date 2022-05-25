package heap

// NewArray 创建类型数组
func (this *Class) NewArray(count uint) *Object {
	if !this.IsArray() {
		panic("Not array class: " + this.name)
	}
	switch this.Name() {
	case "[Z": return &Object{class:this, data: make([]int8, count)}
	case "[B": return &Object{class:this, data: make([]int8, count)}
	case "[C": return &Object{class:this, data: make([]uint16, count)}
	case "[S": return &Object{class:this, data: make([]int16, count)}
	case "[I": return &Object{class:this, data: make([]int32, count)}
	case "[J": return &Object{class:this, data: make([]int64, count)}
	case "[F": return &Object{class:this, data: make([]float32, count)}
	case "[D": return &Object{class:this, data: make([]float64, count)}
	default: return &Object{class:this, data: make([]*Object, count)}
	}
}

