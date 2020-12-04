package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttributeAttribute struct {
	MarkerAttribute
}
type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}
