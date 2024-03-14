package asm

import (
	"fmt"
	"strings"
)

type ValueType byte

const (
	VT_Illigal ValueType = iota
	VT_I64
	VT_Bool
	VT_F64
	VT_String
	VT_List
)

func (vt ValueType) Inspect() string {
	switch vt {
	case VT_I64:
		return "i64"
	case VT_F64:
		return "f64"
	case VT_Bool:
		return "bool"
	case VT_String:
		return "str"
	case VT_List:
		return "list"
	default:
		return "illigal"
	}
}

type Value struct {
	Type ValueType

	Integer64 int64
	Float64   float64
	Boolean   bool
	String    string

	Items []Value
}

func (vl *Value) Inspect() string {
	switch vl.Type {
	case VT_I64:
		return fmt.Sprintf("%d", vl.Integer64)
	case VT_F64:
		return fmt.Sprintf("%f", vl.Float64)
	case VT_Bool:
		return fmt.Sprintf("%v", vl.Boolean)
	case VT_String:
		return fmt.Sprintf("\"%s\"", vl.String)
	case VT_List:
		items := []string{}
		for _, v := range vl.Items {
			items = append(items, v.Inspect())
		}
		return fmt.Sprintf("'(%s)", strings.Join(items, " "))
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

func ValueF64(val float64) Value {
	return Value{
		Type:    VT_F64,
		Float64: val,
	}
}

func ValueBool(val bool) Value {
	return Value{
		Type:    VT_Bool,
		Boolean: val,
	}
}

func ValueString(val string) Value {
	return Value{
		Type:   VT_String,
		String: val,
	}
}

func ValueList(values ...Value) Value {
	return Value{
		Type:  VT_List,
		Items: append([]Value{}, values...),
	}
}
