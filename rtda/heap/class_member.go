package heap

import "jvm-go/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (this *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	this.accessFlags = memberInfo.AccessFlags()
	this.name = memberInfo.Name()
	this.descriptor = memberInfo.Descriptor()
}
func (this *ClassMember) IsStatic() bool {
	return this.accessFlags&ACC_STATIC != 0
}
func (this *ClassMember) IsPublic() bool {
	return this.accessFlags&ACC_PUBLIC != 0
}
func (this *ClassMember) IsFinal() bool {
	return this.accessFlags&ACC_FINAL != 0
}
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (this *ClassMember) IsProtected() bool {
	return this.accessFlags&ACC_PROTECTED != 0

}

func (this *ClassMember) IsPrivate() bool {
	return this.accessFlags&ACC_PRIVATE != 0

}

func (this *ClassMember) Class() *Class {
	return this.class
}
func (this *ClassMember) Name() string {
	return this.name
}
