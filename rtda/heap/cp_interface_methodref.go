package heap

import "jvm-go/classfile"

// InterfaceMethodRef 接口方法引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// ResolvedInterfaceMethod 查找接口方法引用
func (this *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if this.method == nil {
		this.resolveInterfaceMethodRef()
	}
	return this.method
}
func (this *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := this.cp.class
	c := this.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, this.name, this.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.method = method
}
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
