package rtda

import "gvm/src/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
