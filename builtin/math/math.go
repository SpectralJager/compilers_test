package math

import (
	"fmt"
	"grimlang/object"
	"math"
)

func Sqrt(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("root: expect 1 arguments, got %d", len(args))
	}
	if !object.Is(args[0].Kind(), object.FloatLitteral) {
		return nil, fmt.Errorf("root: value should be float, got %s", args[0].Type().Inspect())
	}
	value := *args[0].(*object.LitteralFloat)
	value.Value = math.Sqrt(value.Value)
	return &value, nil
}

func Pow(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("pow: expect 2 arguments, got %d", len(args))
	}
	if !object.Is(args[0].Kind(), object.FloatLitteral) {
		return nil, fmt.Errorf("pow: value should be float, got %s", args[0].Type().Inspect())
	}
	value := *args[0].(*object.LitteralFloat)
	if !object.Is(args[1].Kind(), object.FloatLitteral) {
		return nil, fmt.Errorf("pow: power should be float, got %s", args[1].Type().Inspect())
	}
	power := *args[1].(*object.LitteralFloat)
	value.Value = math.Pow(value.Value, power.Value)
	return &value, nil
}
