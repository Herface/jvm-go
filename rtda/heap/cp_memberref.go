package heap

// 类成员符号引用

import (
	"jvm-go/classfile"
)

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (this *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	this.className = refInfo.ClassName()
	this.name, this.descriptor = refInfo.NameAndDescriptor()
}
func (this *MemberRef) Name() string {
	return this.name
}

func (this *MemberRef) Descriptor() string {
	return this.descriptor
}
