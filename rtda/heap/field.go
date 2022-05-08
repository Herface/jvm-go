package heap

import "jvm-go/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyAttributes(cfField)
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}
func (this *Field) isLongOrDouble() bool {
	return this.descriptor == "J" || this.descriptor == "D"
}

func (this *Field) SlotId() uint {
	return this.slotId
}

func (this *Field) ConstValueIndex() uint {
	return this.constValueIndex
}

func (this *Field) Descriptor() string {
	return this.descriptor
}