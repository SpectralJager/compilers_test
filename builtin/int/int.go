package builtinInt

import (
	"fmt"
	"grimlang/object"
	"strconv"
)

func IntAdd(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("add: expect atleast 2 arguments, got %d", len(args))
	}
	res := 0
	for i, arg := range args {
		if !object.Is(object.IntLitteral, arg.Kind()) {
			return nil, fmt.Errorf("add: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		res += arg.(*object.LitteralInt).Value
	}
	return &object.LitteralInt{Value: res}, nil
}

func IntSub(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("sub: expect atleast 2 arguments, got %d", len(args))
	}
	var res int
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("sub: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			res = arg.(*object.LitteralInt).Value
			continue
		}
		res -= arg.(*object.LitteralInt).Value
	}
	return &object.LitteralInt{Value: res}, nil
}

func IntMul(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("mul: expect atleast 2 arguments, got %d", len(args))
	}
	var res int
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("mul: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			res = arg.(*object.LitteralInt).Value
			continue
		}
		if arg.(*object.LitteralInt).Value == 0 {
			return &object.LitteralInt{Value: 0}, nil
		}
		res *= arg.(*object.LitteralInt).Value
	}
	return &object.LitteralInt{Value: res}, nil
}

func IntDiv(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("div: expect atleast 2 arguments, got %d", len(args))
	}
	var res int
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("div: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			res = arg.(*object.LitteralInt).Value
			continue
		}
		if arg.(*object.LitteralInt).Value == 0 {
			return nil, fmt.Errorf("div: division by zero")
		}
		res /= arg.(*object.LitteralInt).Value
	}
	return &object.LitteralInt{Value: res}, nil
}

func IntLt(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("lt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralInt
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("lt: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralInt)
			continue
		}
		if target.Value >= arg.(*object.LitteralInt).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralInt)
	}
	return &object.LitteralBool{Value: true}, nil
}

func IntGt(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("gt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralInt
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("gt: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralInt)
			continue
		}
		if target.Value <= arg.(*object.LitteralInt).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralInt)
	}
	return &object.LitteralBool{Value: true}, nil
}

func IntLeq(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("leq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralInt
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("leq: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralInt)
			continue
		}
		if target.Value > arg.(*object.LitteralInt).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralInt)
	}
	return &object.LitteralBool{Value: true}, nil
}

func IntGeq(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("geq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralInt
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("geq: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralInt)
			continue
		}
		if target.Value < arg.(*object.LitteralInt).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralInt)
	}
	return &object.LitteralBool{Value: true}, nil
}

func IntEq(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("eq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralInt
	for i, arg := range args {
		if !object.Is(arg.Kind(), object.IntLitteral) {
			return nil, fmt.Errorf("eq: argument #%d should be int, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralInt)
			continue
		}
		if target.Value != arg.(*object.LitteralInt).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralInt)
	}
	return &object.LitteralBool{Value: true}, nil
}

func IntToString(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toString: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !object.Is(arg.Kind(), object.IntLitteral) {
		return nil, fmt.Errorf("toString: argument should be int, got %s", arg.Type().Inspect())
	}
	return &object.LitteralString{Value: strconv.Itoa(arg.(*object.LitteralInt).Value)}, nil
}

func IntToFloat(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toFloat: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !object.Is(arg.Kind(), object.IntLitteral) {
		return nil, fmt.Errorf("toFloat: argument should be int, got %s", arg.Type().Inspect())
	}
	return &object.LitteralFloat{Value: float64(arg.(*object.LitteralInt).Value)}, nil
}

func IntToBool(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toBool: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !object.Is(arg.Kind(), object.IntLitteral) {
		return nil, fmt.Errorf("toBool: argument should be int, got %s", arg.Type().Inspect())
	}
	return &object.LitteralBool{Value: arg.(*object.LitteralInt).Value != 0}, nil
}
