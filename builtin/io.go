package builtin

import (
	"fmt"
	"grimlang/dtype"
	"grimlang/object"
	"strings"
)

func IoPrintln(args ...object.Object) (object.Object, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("println: expect atleast 1 argument, got %d", len(args))
	}
	vals := []string{}
	for i, arg := range args {
		if !new(dtype.StringType).Compare(arg.Type()) {
			return nil, fmt.Errorf("println: argument #%d should be string, got %s", i, arg.Type().Name())
		}
		vals = append(vals, arg.(*object.StringObject).Value)
	}
	fmt.Println(strings.Join(vals, " "))
	return nil, nil
}
