package ir

import (
	"fmt"
)

type Instruction interface {
	fmt.Stringer
	instr()
}
