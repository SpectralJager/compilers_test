package runtime

import (
	"fmt"
	"grimlang/backend/asm"
	"log"
)

func Must[T any](val T, err error) T {
	if err != nil {
		log.Fatalf("something goes wrong during execution -> %s", err.Error())
	}
	return val
}

func I64Value(val asm.Value) int64 {
	if val.Type != asm.VT_I64 {
		panic(fmt.Errorf("can't get %s from %s", asm.VT_I64.Inspect(), val.Type.Inspect()))
	}
	return val.Integer64
}

func BoolValue(val asm.Value) bool {
	if val.Type != asm.VT_Bool {
		panic(fmt.Errorf("can't get %s from %s", asm.VT_I64.Inspect(), val.Type.Inspect()))
	}
	return val.Boolean
}
