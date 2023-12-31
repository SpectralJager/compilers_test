package builtin

import (
	"fmt"
	"grimlang/object"
	"os"
)

func Exit(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("exit: expect 1 argument, got %d", len(args))
	}
	obj := args[0]
	if !new(object.DTypeInt).Compare(obj.Type()) {
		return nil, fmt.Errorf("exit: exit code should be int, got %s", obj.Type().Inspect())
	}
	code := args[0].(*object.LitteralInt)
	os.Exit(code.Value)
	return nil, nil
}
