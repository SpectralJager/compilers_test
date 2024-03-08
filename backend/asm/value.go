package asm

import "fmt"

type ValueType byte

const (
	VT_Illigal ValueType = iota
	VT_I64
)

func (vt ValueType) Inspect() string {
	switch vt {
	case VT_I64:
		return "(i64)"
	default:
		return "(illigal)"
	}
}

type Value struct {
	Type      ValueType
	Integer64 int64
}

func (vl *Value) Inspect() string {
	switch vl.Type {
	case VT_I64:
		return fmt.Sprintf("%d", vl.Integer64)
	default:
		return "illigal"
	}
}

func ValueI64(val int64) Value {
	return Value{
		Type:      VT_I64,
		Integer64: val,
	}
}
