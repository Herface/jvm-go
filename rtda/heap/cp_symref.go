package heap

// SymRef 符号引用
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (this *SymRef) ResolvedClass() *Class {
	if this.class == nil {
		this.resolveClassRef()
	}
	return this.class
}
func (this *SymRef) resolveClassRef() {
	d := this.cp.class
	c := d.loader.LoadClass(this.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.class = c
}
