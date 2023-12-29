package builtin

import (
	"fmt"
	"grimlang/dtype"
	"grimlang/object"
	"strconv"
)

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

func IntSub(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("sub: expect atleast 2 arguments, got %d", len(args))
	}
	var res int
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("sub: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			res = arg.(*object.IntObject).Value
			continue
		}
		res -= arg.(*object.IntObject).Value
	}
	return &object.IntObject{Value: res}, nil
}

func IntMul(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("mul: expect atleast 2 arguments, got %d", len(args))
	}
	var res int
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("mul: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			res = arg.(*object.IntObject).Value
			continue
		}
		if arg.(*object.IntObject).Value == 0 {
			return &object.IntObject{Value: 0}, nil
		}
		res *= arg.(*object.IntObject).Value
	}
	return &object.IntObject{Value: res}, nil
}

func IntDiv(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("div: expect atleast 2 arguments, got %d", len(args))
	}
	var res int
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("div: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			res = arg.(*object.IntObject).Value
			continue
		}
		if arg.(*object.IntObject).Value == 0 {
			return nil, fmt.Errorf("div: division by zero")
		}
		res /= arg.(*object.IntObject).Value
	}
	return &object.IntObject{Value: res}, nil
}

func IntLt(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("lt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.IntObject
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("lt: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.IntObject)
			continue
		}
		if target.Value >= arg.(*object.IntObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func IntGt(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("gt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.IntObject
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("gt: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.IntObject)
			continue
		}
		if target.Value <= arg.(*object.IntObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func IntLeq(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("leq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.IntObject
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("leq: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.IntObject)
			continue
		}
		if target.Value > arg.(*object.IntObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func IntGeq(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("geq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.IntObject
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("geq: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.IntObject)
			continue
		}
		if target.Value < arg.(*object.IntObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func IntEq(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("eq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.IntObject
	for i, arg := range args {
		if !new(dtype.IntType).Compare(arg.Type()) {
			return nil, fmt.Errorf("eq: argument #%d should be int, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.IntObject)
			continue
		}
		if target.Value != arg.(*object.IntObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func IntToString(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toString: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(dtype.IntType).Compare(arg.Type()) {
		return nil, fmt.Errorf("toString: argument should be int, got %s", arg.Type().Name())
	}
	return &object.StringObject{Value: strconv.Itoa(arg.(*object.IntObject).Value)}, nil
}

func IntToFloat(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toFloat: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(dtype.IntType).Compare(arg.Type()) {
		return nil, fmt.Errorf("toFloat: argument should be int, got %s", arg.Type().Name())
	}
	return &object.FloatObject{Value: float64(arg.(*object.IntObject).Value)}, nil
}

func IntToBool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toBool: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(dtype.IntType).Compare(arg.Type()) {
		return nil, fmt.Errorf("toBool: argument should be int, got %s", arg.Type().Name())
	}
	return &object.BoolObject{Value: arg.(*object.IntObject).Value != 0}, nil
}
