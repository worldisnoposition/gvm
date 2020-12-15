package base

import (
	"gvm/src/rtda"
	"gvm/src/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	inistSuperClass(thread, class)
}

func inistSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}
