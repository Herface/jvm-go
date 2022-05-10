package heap

import "jvm-go/classfile"

// Method 类方法信息
type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()

	}
	return methods
}
func (this *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		this.maxStack = uint(codeAttr.MaxStack())
		this.maxLocals = uint(codeAttr.MaxLocals())
		this.code = codeAttr.Code()
	}
}
func (this *Method) MaxStack() uint {
	return this.maxStack
}
func (this *Method) MaxLocals() uint {
	return this.maxLocals
}

func (this *Method) Code() []byte {
	return this.code
}

func (this *Method) ArgSlotCount() uint {
	return this.argSlotCount
}

func (this *Method) calcArgSlotCount() {

	parsedDescriptor := parseMethodDescriptor(this.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		this.argSlotCount++
		if paramType == "J" || paramType == "D" {
			this.argSlotCount++
		}
	}
	if !this.IsStatic() {
		this.argSlotCount++
	}

}
