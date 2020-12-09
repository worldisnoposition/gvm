package heap

import "gvm/src/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}

}

func (self *Field) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (self *Field) Class() *Class {
	return self.class
}

func (self *Field) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *Field) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Field) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *Field) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *Field) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Field) Descriptor() string {
	return self.descriptor
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
