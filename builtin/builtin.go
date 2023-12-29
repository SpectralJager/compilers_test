package builtin

import (
	"fmt"
	"grimlang/dtype"
	"grimlang/object"
	"os"
)

type BuiltinFunction func(args ...object.Object) (object.Object, error)

// =========================
func Exit(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("exit: expect 1 argument, got %d", len(args))
	}
	obj := args[0]
	if !new(dtype.IntType).Compare(obj.Type()) {
		return nil, fmt.Errorf("exit: exit code should be int, got %s", obj.Type().Name())
	}
	code := args[0].(*object.IntObject)
	os.Exit(code.Value)
	return nil, nil
}
