package heap

import (
	"jvm-go/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string // thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (this *Class) IsPublic() bool {
	return 0 != this.accessFlags&ACC_PUBLIC
}
func (this *Class) isAccessibleTo(other *Class) bool {
	return this.IsPublic() || this.getPackageName() == other.getPackageName()
}
func (this *Class) getPackageName() string {
	if i := strings.LastIndex(this.name, "/"); i >= 0 {
		return this.name[:i]
	}
	return ""
}
func (this *Class) isSubClassOf(other *Class) bool {
	for c := this.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (this *Class) ConstantPool() *ConstantPool {
	return this.constantPool
}

func (this *Class) IsInterface() bool {
	return this.accessFlags&ACC_INTERFACE != 0
}

func (this *Class) IsAbstract() bool {
	return this.accessFlags&ACC_ABSTRACT != 0
}

func (this *Class) NewObject() *Object {
	return newObject(this)
}

func (this *Class) StaticVars() Slots {
	return this.staticVars
}

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}
func (this *Class) isImplements(iface *Class) bool {
	for c := this; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// isSubInterfaceOf 由下至上 匹配父接口
func (this *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range this.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}
