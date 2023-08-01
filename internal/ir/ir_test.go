package ir

import (
	"fmt"
	"testing"
)

func TestIR(t *testing.T) {
	ir := Program{
		Name: "test",
		Constants: []IConstant{
			&Integer{40},
			&Integer{2},
			&Integer{1},
		},
		InitCode: NewCode().
			WriteBytes(Load(0)...).
			WriteByte(Halt()),
	}
	fmt.Println(ir.String())
}
