package builtinFloat

import (
	"fmt"
	"grimlang/object"
	"strconv"
)

func FloatAdd(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("add: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("add: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		res += arg.(*object.LitteralFloat).Value
	}
	return &object.LitteralFloat{Value: res}, nil
}

func FloatSub(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("sub: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("sub: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			res = arg.(*object.LitteralFloat).Value
			continue
		}
		res -= arg.(*object.LitteralFloat).Value
	}
	return &object.LitteralFloat{Value: res}, nil
}

func FloatMul(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("mul: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("mul: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			res = arg.(*object.LitteralFloat).Value
			continue
		}
		if arg.(*object.LitteralFloat).Value == 0 {
			return &object.LitteralFloat{Value: 0}, nil
		}
		res *= arg.(*object.LitteralFloat).Value
	}
	return &object.LitteralFloat{Value: res}, nil
}

func FloatDiv(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("div: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("div: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			res = arg.(*object.LitteralFloat).Value
			continue
		}
		if arg.(*object.LitteralFloat).Value == 0 {
			return nil, fmt.Errorf("div: division by zero")
		}
		res /= arg.(*object.LitteralFloat).Value
	}
	return &object.LitteralFloat{Value: res}, nil
}

func FloatLt(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("lt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralFloat
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("lt: argument #%d should be Float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralFloat)
			continue
		}
		if target.Value >= arg.(*object.LitteralFloat).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralFloat)
	}
	return &object.LitteralBool{Value: true}, nil
}

func FloatGt(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("gt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralFloat
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("gt: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralFloat)
			continue
		}
		if target.Value <= arg.(*object.LitteralFloat).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralFloat)
	}
	return &object.LitteralBool{Value: true}, nil
}

func FloatLeq(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("leq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralFloat
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("leq: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralFloat)
			continue
		}
		if target.Value > arg.(*object.LitteralFloat).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralFloat)
	}
	return &object.LitteralBool{Value: true}, nil
}

func FloatGeq(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("geq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralFloat
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("geq: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralFloat)
			continue
		}
		if target.Value < arg.(*object.LitteralFloat).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralFloat)
	}
	return &object.LitteralBool{Value: true}, nil
}

func FloatEq(args ...object.Litteral) (object.Litteral, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("eq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.LitteralFloat
	for i, arg := range args {
		if !new(object.DTypeFloat).Compare(arg.Type()) {
			return nil, fmt.Errorf("eq: argument #%d should be float, got %s", i, arg.Type().Inspect())
		}
		if i == 0 {
			target = arg.(*object.LitteralFloat)
			continue
		}
		if target.Value != arg.(*object.LitteralFloat).Value {
			return &object.LitteralBool{Value: false}, nil
		}
		target = arg.(*object.LitteralFloat)
	}
	return &object.LitteralBool{Value: true}, nil
}

func FloatToString(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toString: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(object.DTypeFloat).Compare(arg.Type()) {
		return nil, fmt.Errorf("toString: argument should be Float, got %s", arg.Type().Inspect())
	}
	return &object.LitteralString{Value: strconv.FormatFloat(arg.(*object.LitteralFloat).Value, 'f', -1, 64)}, nil
}

func FloatToInt(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toInt: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(object.DTypeFloat).Compare(arg.Type()) {
		return nil, fmt.Errorf("toInt: argument should be Float, got %s", arg.Type().Inspect())
	}
	return &object.LitteralInt{Value: int(arg.(*object.LitteralFloat).Value)}, nil
}

func FloatToBool(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toBool: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(object.DTypeFloat).Compare(arg.Type()) {
		return nil, fmt.Errorf("toBool: argument should be Float, got %s", arg.Type().Inspect())
	}
	return &object.LitteralBool{Value: arg.(*object.LitteralFloat).Value != 0.0}, nil
}
