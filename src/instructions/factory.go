package instructions

import (
	"fmt"
	"gvm/src/instructions/base"
)
import . "gvm/src/instructions/constants"

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &NOP{}
	case 0x01:
		return &ACONST_NULL{}
		//todo ...还要很多，照github抄下
	default:
		panic(fmt.Errorf("Unsupported opcode:0x%x", opcode))
	}

}
