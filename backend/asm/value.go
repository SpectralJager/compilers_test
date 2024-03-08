package asm

import (
	"errors"
	"fmt"
)

type ValueType byte

const (
	VT_Illigal ValueType = iota
	VT_I64
	VT_Bool

	VT_Symbol
)

func (vt ValueType) Inspect() string {
	switch vt {
	case VT_Symbol:
		return ""
	case VT_I64:
		return "i64"
	case VT_Bool:
		return "bool"
	default:
		return "illigal"
	}
}

type Value struct {
	Type      ValueType
	Ident     string
	Integer64 int64
	Boolean   bool
}

func (vl *Value) Inspect() string {
	switch vl.Type {
	case VT_Symbol:
		return vl.Ident
	case VT_I64:
		return fmt.Sprintf("%d", vl.Integer64)
	case VT_Bool:
		return fmt.Sprintf("%v", vl.Boolean)
	default:
		return "illigal"
	}
}

func (vl *Value) Compare(other Value) error {
	if vl.Type == VT_Symbol {
		return errors.New("can't compare symbol value")
	}
	if vl.Type != other.Type {
		return fmt.Errorf("%s <> %s", vl.Type.Inspect(), other.Type.Inspect())
	}
	return nil
}

func ValueSymbol(ident string) Value {
	return Value{
		Type:  VT_Symbol,
		Ident: ident,
	}
}

func ValueI64(val int64) Value {
	return Value{
		Type:      VT_I64,
		Integer64: val,
	}
}

func ValueBool(val bool) Value {
	return Value{
		Type:    VT_Bool,
		Boolean: val,
	}
}
