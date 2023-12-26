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

func IntAdd(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("add: expect atleast 2 arguments, got %d", len(args))
	}
	res := 0
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("add: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		res += arg.(*object.IntObject).Value
	}
	return &object.IntObject{Value: res}, nil
}
