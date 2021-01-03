package heap

type Object struct {
	class *Class
	//fields Slots
	data  interface{}
	extra interface{}
}

func (self *Object) Data() interface{} {
	return self.data
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
