package heap

import (
	"jvm-go/classfile"
	"strings"
)

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// Class class的运行时表示 非java.lang.Class
type Class struct {
	accessFlags    uint16
	name           string // thisClassName
	superClassName string
	interfaceNames []string
	constantPool   *ConstantPool
	fields         []*Field
	methods        []*Method
	loader         *ClassLoader
	superClass     *Class
	interfaces     []*Class
	// 实例属性数量
	instanceSlotCount uint
	// 静态属性数量
	staticSlotCount uint
	staticVars      Slots
	initStarted     bool
	jClass          *Object
	sourceFile      string
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
	class.sourceFile = getSourceFile(cf)
	return class
}
func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}
func (this *Class) IsPublic() bool {
	return 0 != this.accessFlags&ACC_PUBLIC
}
func (this *Class) isAccessibleTo(other *Class) bool {
	return this.IsPublic() || this.GetPackageName() == other.GetPackageName()
}
func (this *Class) GetPackageName() string {
	if i := strings.LastIndex(this.name, "/"); i >= 0 {
		return this.name[:i]
	}
	return ""
}
func (this *Class) IsSubClassOf(other *Class) bool {
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

func (this *Class) IsImplements(iface *Class) bool {
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

func (this *Class) IsSuperClassOf(cls *Class) bool {
	superClass := cls.superClass
	if superClass == nil {
		return false
	}
	if superClass == this {
		return true
	}
	if this.IsSuperClassOf(superClass) {
		return true
	}
	return false
}

func (this *Class) SuperClass() *Class {
	return this.superClass
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (this *Class) FindMainClass() *Method {
	for _, m := range this.methods {
		if m.name == "main" && m.descriptor == "()V" {
			return m
		}
	}
	return nil
}
func (this *Class) InitStarted() bool {
	return this.initStarted
}
func (this *Class) StartInit() {
	this.initStarted = true
}

func (this *Class) GetClinitMethod() *Method {
	for _, m := range this.methods {
		if m.name == "<clinit>" && m.descriptor == "()V" {
			return m
		}
	}
	return nil
}

func (this *Class) Name() string {
	return this.name
}

func (this *Class) IsArray() bool {
	return strings.HasPrefix(this.name, "[")
}

func (this *Class) Loader() *ClassLoader {
	return this.loader
}

// ArrayClass 加载引用类型的数组类型
func (this *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(this.name)
	return this.loader.LoadClass(arrayClassName)
}
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}
func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}

func (this *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(this.name)
	return this.loader.LoadClass(componentClassName)
}

// getComponentClassName 数组元素类型
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}
func toClassName(descriptor string) string {
	if descriptor[0] == '[' { // array
		return descriptor
	}
	if descriptor[0] == 'L' { // object
		return descriptor[1 : len(descriptor)-1]
	}
	for className, d := range primitiveTypes {
		if d == descriptor { // primitive
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
func (this *Class) isAssignableFrom(other *Class) bool {
	s, t := other, this
	if s == t {
		return true
	}
	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

}

func (this *Class) isJlObject() bool {
	return this.name == "java.lang.Object"
}

func (this *Class) isJlCloneable() bool {
	return this.name == "java.lang.Cloneable"
}

func (this *Class) isJioSerializable() bool {
	return this.name == "java.lang.Serializable"

}

func (this *Class) isSuperInterfaceOf(s *Class) bool {
	for _, intf := range s.interfaces {
		if intf == this {
			return true
		}
	}
	return false
}

func (this *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := this; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (this *Class) JClass() *Object {
	return this.jClass
}

func (this *Class) JavaName() string {
	return strings.Replace(this.name, "/", ".", -1)
}

func (this *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[this.name]
	return ok

}

func (this *Class) SourceFile() string {
	return this.sourceFile
}

func (this *Class) GetMethod(name, descriptor string) *Method {
	for _, m := range this.methods {
		if m.name == name && m.descriptor == descriptor {
			return m
		}
	}
	return nil
}
