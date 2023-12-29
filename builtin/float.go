package builtin

import (
	"fmt"
	"grimlang/dtype"
	"grimlang/object"
	"strconv"
)

func FloatAdd(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("add: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("add: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		res += arg.(*object.FloatObject).Value
	}
	return &object.FloatObject{Value: res}, nil
}

func FloatSub(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("sub: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("sub: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			res = arg.(*object.FloatObject).Value
			continue
		}
		res -= arg.(*object.FloatObject).Value
	}
	return &object.FloatObject{Value: res}, nil
}

func FloatMul(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("mul: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("mul: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			res = arg.(*object.FloatObject).Value
			continue
		}
		if arg.(*object.FloatObject).Value == 0 {
			return &object.FloatObject{Value: 0}, nil
		}
		res *= arg.(*object.FloatObject).Value
	}
	return &object.FloatObject{Value: res}, nil
}

func FloatDiv(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("div: expect atleast 2 arguments, got %d", len(args))
	}
	var res float64
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("div: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			res = arg.(*object.FloatObject).Value
			continue
		}
		if arg.(*object.FloatObject).Value == 0 {
			return nil, fmt.Errorf("div: division by zero")
		}
		res /= arg.(*object.FloatObject).Value
	}
	return &object.FloatObject{Value: res}, nil
}

func FloatLt(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("lt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.FloatObject
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("lt: argument #%d should be Float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.FloatObject)
			continue
		}
		if target.Value >= arg.(*object.FloatObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func FloatGt(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("gt: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.FloatObject
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("gt: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.FloatObject)
			continue
		}
		if target.Value <= arg.(*object.FloatObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func FloatLeq(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("leq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.FloatObject
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("leq: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.FloatObject)
			continue
		}
		if target.Value > arg.(*object.FloatObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func FloatGeq(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("geq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.FloatObject
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("geq: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.FloatObject)
			continue
		}
		if target.Value < arg.(*object.FloatObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func FloatEq(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("eq: expect atleast 2 arguments, got %d", len(args))
	}
	var target *object.FloatObject
	for i, arg := range args {
		if !new(dtype.FloatType).Compare(arg.Type()) {
			return nil, fmt.Errorf("eq: argument #%d should be float, got %s", i, arg.Type().Name())
		}
		if i == 0 {
			target = arg.(*object.FloatObject)
			continue
		}
		if target.Value != arg.(*object.FloatObject).Value {
			return &object.BoolObject{Value: false}, nil
		}
	}
	return &object.BoolObject{Value: true}, nil
}

func FloatToString(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toString: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(dtype.FloatType).Compare(arg.Type()) {
		return nil, fmt.Errorf("toString: argument should be Float, got %s", arg.Type().Name())
	}
	return &object.StringObject{Value: strconv.FormatFloat(arg.(*object.FloatObject).Value, 'f', -1, 64)}, nil
}

func FloatToInt(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toInt: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(dtype.FloatType).Compare(arg.Type()) {
		return nil, fmt.Errorf("toInt: argument should be Float, got %s", arg.Type().Name())
	}
	return &object.IntObject{Value: int(arg.(*object.FloatObject).Value)}, nil
}

func FloatToBool(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("toBool: expect 1 argument, got %d", len(args))
	}
	arg := args[0]
	if !new(dtype.FloatType).Compare(arg.Type()) {
		return nil, fmt.Errorf("toBool: argument should be Float, got %s", arg.Type().Name())
	}
	return &object.BoolObject{Value: arg.(*object.FloatObject).Value != 0.0}, nil
}
