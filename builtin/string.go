package builtin

import (
	"fmt"
	"grimlang/dtype"
	"grimlang/object"
)

func StringConcat(args ...object.Object) (object.Object, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("concat: expect atleast 1 argument, got %d", len(args))
	}
	res := ""
	for i, arg := range args {
		if !new(dtype.StringType).Compare(arg.Type()) {
			return nil, fmt.Errorf("println: argument #%d should be string, got %s", i, arg.Type().Name())
		}
		res += arg.(*object.StringObject).Value
	}
	return &object.StringObject{Value: res}, nil
}
