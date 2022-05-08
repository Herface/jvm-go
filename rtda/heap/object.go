package heap

type Object struct {
	class  *Class
	fields Slots
}

func (this *Object) Fields() Slots {
	return this.fields
}
func (this *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(this.class)
}
