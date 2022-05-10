package factory

import (
	"fmt"
	"jvm-go/instructions/common"
	"jvm-go/instructions/comparisons"
	"jvm-go/instructions/constants"
	"jvm-go/instructions/control"
	"jvm-go/instructions/conversions"
	"jvm-go/instructions/loads"
	"jvm-go/instructions/math"
	"jvm-go/instructions/references"
	"jvm-go/instructions/stores"
)

func NewInstruction(opcode byte) common.Instruction {

	switch opcode {
	case 0x0:
		return &common.NoOperandsInstruction{}
	case 0x1:
		return &constants.ACONST_NULL{}
	case 0x19:
		return &loads.ALOAD{}
	case 0x2a:
		return &loads.ALOAD_0{}
	case 0x2b:
		return &loads.ALOAD_1{}
	case 0x2c:
		return &loads.ALOAD_2{}
	case 0x2d:
		return &loads.ALOAD_3{}
	case 0x3a:
		return &stores.ASTORE{}
	case 0x4b:
		return &stores.ASTORE_0{}
	case 0x4c:
		return &stores.ASTORE_1{}
	case 0x4d:
		return &stores.ASTORE_2{}
	case 0x4e:
		return &stores.ASTORE_3{}
	case 0x10:
		return &constants.BIPUSH{}
	case 0x90:
		return &conversions.D2F{}
	case 0x8e:
		return &conversions.D2I{}
	case 0x8f:
		return &conversions.D2L{}
	case 0x63:
		return &math.DADD{}
	case 0x98:
		return &comparisons.DCMPG{}
	case 0x97:
		return &comparisons.DCMPL{}
	case 0xe:
		return &constants.DCONST_0{}
	case 0xf:
		return &constants.DCONST_1{}

	case 0xB1:
		return &control.RETURN{}

	case 0xB8:
		return &references.INVOKE_STATIC{}

	}

	panic(fmt.Sprintf("unknown code: %v", opcode))
}
