package builtinList

import (
	"fmt"
	"grimlang/object"
)

func Len(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("len: expect 1 argument, got %d", len(args))
	}
	lst := args[0]
	if !object.Is(lst.Kind(), object.ListLitteral) {
		return nil, fmt.Errorf("len: argument should be list, got %s", lst.Type().Inspect())
	}
	return &object.LitteralInt{
		Value: len(lst.(*object.LitteralList).Items),
	}, nil
}

func Get(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("get: expect 2 argument, got %d", len(args))
	}
	lst := args[0]
	if !object.Is(lst.Kind(), object.ListLitteral) {
		return nil, fmt.Errorf("get: 1th argument should be list, got %s", lst.Type().Inspect())
	}
	index := args[1]
	if !object.Is(index.Kind(), object.IntLitteral) {
		return nil, fmt.Errorf("get: 2th argument should be int, got %s", lst.Type().Inspect())
	}

	return lst.(*object.LitteralList).Items[index.(*object.LitteralInt).Value], nil
}

func Set(args ...object.Litteral) (object.Litteral, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("set: expect 3 argument, got %d", len(args))
	}
	lst := args[0]
	if !object.Is(lst.Kind(), object.ListLitteral) {
		return nil, fmt.Errorf("set: 1th argument should be list, got %s", lst.Type().Inspect())
	}
	index := args[1]
	if !object.Is(index.Kind(), object.IntLitteral) {
		return nil, fmt.Errorf("set: 2th argument should be int, got %s", lst.Type().Inspect())
	}
	item := lst.(*object.LitteralList).
		Items[index.(*object.LitteralInt).Value]
	val := args[2]
	if !item.Type().Compare(val.Type()) {
		return nil, fmt.Errorf("set: item and value types mismatch: %s != %s", item.Type().Inspect(), val.Type().Inspect())
	}
	lst.(*object.LitteralList).
		Items[index.(*object.LitteralInt).Value] = val
	return nil, nil
}
