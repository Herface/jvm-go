package heap

import (
	"fmt"
	"jvm-go/classfile"
	"jvm-go/classpath"
)

// ClassLoader 类加载器
type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath,) *ClassLoader {
	loader := &ClassLoader{
		cp: cp,
		//verboseFlag: verboseFlag,
		classMap: make(map[string]*Class),
	}
	// 先加载java.lang.Class
	loader.loadBasicClasses()
	// 基本数据类型
	loader.loadPrimitiveClasses()
	return loader
}
func (this *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		this.loadPrimitiveClass(primitiveType)
	}
}

func (this *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		name: className,
		loader: this,
		initStarted: true,
	}
	class.jClass = this.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	this.classMap[className] = class
}

func (this *ClassLoader) loadBasicClasses() {
	jlClassClass := this.LoadClass("java/lang/Class")
	for _, class := range this.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

func (this *ClassLoader) LoadClass(name string) *Class {
	if class, ok := this.classMap[name]; ok {
		return class // already loaded
	}
	var class *Class
	if name[0] == '[' { // array class
		class = this.loadArrayClass(name)
	} else {
		class = this.loadNonArrayClass(name)
	}
	// 给加载的类(go结构体表示) 附上java.lang.Class对象
	if jlClassClass, ok := this.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}
	return class
}

func (this *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		name: name,
		loader: this,
		initStarted: true,
		superClass: this.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			this.LoadClass("java/lang/Cloneable"),
			this.LoadClass("java/io/Serializable"),
		},
	}
	this.classMap[name] = class
	return class
}
func (this *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := this.readClass(name)
	class := this.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}
func (this *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := this.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}
func (this *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = this
	resolveSuperClass(class)
	resolveInterfaces(class)
	this.classMap[class.name] = class
	return class
}
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 链接
func link(class *Class) {
	verify(class)
	prepare(class)
	resolve(class)
}
func verify(class *Class) {
	// todo
}

// resolve 解析 符号引用 -> 直接引用
// 静态 private init方法 and 属性
func resolve(class *Class)  {

}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// calcInstanceFieldSlotIds 计算属性所占空间大小
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// 静态属性数量
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 静态属性
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 初始化所有静态常量 <init方法之前>
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			// 将字符串从常量池取出
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}
