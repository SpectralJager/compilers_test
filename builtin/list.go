package builtin

import (
	"fmt"
	"grimlang/dtype"
	"grimlang/object"
)

func ListLen(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("len: expect 1 argument, got %d", len(args))
	}
	lst := args[0]
	if lst.Type().Kind() != dtype.List {
		return nil, fmt.Errorf("len: argument should be list, got %s", lst.Type().Name())
	}
	return &object.IntObject{
		Value: len(lst.(*object.ListObject).Items),
	}, nil
}

func ListGet(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("get: expect 2 argument, got %d", len(args))
	}
	lst := args[0]
	if lst.Type().Kind() != dtype.List {
		return nil, fmt.Errorf("get: 1th argument should be list, got %s", lst.Type().Name())
	}
	index := args[1]
	if index.Type().Kind() != dtype.Int {
		return nil, fmt.Errorf("get: 2th argument should be int, got %s", lst.Type().Name())
	}

	return lst.(*object.ListObject).Items[index.(*object.IntObject).Value], nil
}

func ListSet(args ...object.Object) (object.Object, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("get: expect 3 argument, got %d", len(args))
	}
	lst := args[0]
	if lst.Type().Kind() != dtype.List {
		return nil, fmt.Errorf("get: 1th argument should be list, got %s", lst.Type().Name())
	}
	index := args[1]
	if index.Type().Kind() != dtype.Int {
		return nil, fmt.Errorf("get: 2th argument should be int, got %s", lst.Type().Name())
	}
	item := lst.(*object.ListObject).
		Items[index.(*object.IntObject).Value]
	val := args[2]
	if !item.Type().Compare(val.Type()) {
		return nil, fmt.Errorf("set: item and value types mismatch: %s != %s", item.Type().Name(), val.Type().Name())
	}
	lst.(*object.ListObject).
		Items[index.(*object.IntObject).Value] = val
	return nil, nil
}
