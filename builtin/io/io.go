package builtinIo

import (
	"fmt"
	"grimlang/object"
	"strings"
)

func Println(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("println: expect atleast 1 argument, got %d", len(args))
	}
	vals := []string{}
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.StringLitteral) {
			return nil, fmt.Errorf("println: argument #%d should be string, got %s", i, arg.Type().Inspect())
		}
		vals = append(vals, arg.(*object.LitteralString).Value)
	}
	fmt.Printf("%s\n", strings.Join(vals, " "))
	return nil, nil
}
