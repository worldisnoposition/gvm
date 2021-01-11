package heap

import (
	"gvm/src/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
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
	initStarted       bool
	jClass            *Object
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

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *Class) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) isJ1Object() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJ1Cloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) NewObject() *Object {
	return newObejct(self)
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) isInterface() bool {
	return len(self.interfaces) > 0
}

func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (self *Class) GetMainMthod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name string, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (self *Class) IsSuperClassOf(class *Class) bool {
	return self.superClassName == class.name
}

func (self *Class) GetPackageName() string {
	return self.getPackageName()
}

func (self *Class) IsSubClassOf(class *Class) bool {
	return self.name == class.superClassName
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER

}

func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) Name() string {
	return self.name
}

func newObejct(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}
func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

func getComponentClassName(className string) string {
	if className[0] == '[' {
		compnentTypeDescriptor := className[1:]
		return toClassName(compnentTypeDescriptor)
	}
	panic("Not array: " + className)
}

func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
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

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", 1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

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
