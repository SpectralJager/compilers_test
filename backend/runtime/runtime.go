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

func SymbolValue(val asm.Value) (string, error) {
	if val.Type != asm.VT_Symbol {
		return "", fmt.Errorf("can't get %s from %s", asm.VT_I64.Inspect(), val.Type.Inspect())
	}
	return val.Ident, nil
}

func I64Value(val asm.Value) (int64, error) {
	if val.Type != asm.VT_I64 {
		return 0, fmt.Errorf("can't get %s from %s", asm.VT_I64.Inspect(), val.Type.Inspect())
	}
	return val.Integer64, nil
}

func BoolValue(val asm.Value) (bool, error) {
	if val.Type != asm.VT_Bool {
		return false, fmt.Errorf("can't get %s from %s", asm.VT_I64.Inspect(), val.Type.Inspect())
	}
	return val.Boolean, nil
}
