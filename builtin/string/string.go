package builtinString

import (
	"fmt"
	"grimlang/object"
	"strings"
)

func Format(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("format: expect atleast 1 arguments, got %d", len(args))
	}
	if !object.Is(args[0].Kind(), object.StringLitteral) {
		return nil, fmt.Errorf("format: format argument should be string, got %s", args[0].Type().Inspect())
	}
	format := *args[0].(*object.LitteralString)
	c := strings.Count(format.Value, "$$")
	if c == 0 {
		return &format, nil
	}
	if c != len(args)-1 {
		return nil, fmt.Errorf("format: contains %d entries, got %d arguments", c, len(args)-1)
	}
	for i, arg := range args[1:] {
		if !object.Is(args[0].Kind(), object.StringLitteral) {
			return nil, fmt.Errorf("format: argument #%d should be string, got %s", i, arg.Type().Inspect())
		}
		argString := *arg.(*object.LitteralString)
		format.Value = strings.Replace(format.Value, "$$", argString.Value, 1)
	}
	return &format, nil
}
