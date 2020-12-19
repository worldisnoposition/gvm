package heap

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.isInterface() {
				return s.isSubClassOf(t)
			} else {
				return s.isImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return s.isJ1Object()
			} else {
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJ1Object()
			} else {
				return t.isJ1Cloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}
	return false
	//if !t.isInterface() {
	//	return s.isSubClassOf(t)
	//} else {
	//	return s.isImplements(t)
	//}
}
