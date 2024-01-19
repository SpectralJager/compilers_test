package runtime

import (
	"fmt"
	"strings"
)

type _type struct {
	kind Kind
	name string

	subType Type

	args []Type
	ret  Type

	fields []Field
}

func (typ _type) Kind() Kind {
	return typ.kind
}

func (typ _type) Name() string {
	switch typ.kind {
	case Int, Float, String, Bool, Any, Void, List, Function, Null:
		return typ.kind.String()
	case Record:
		return typ.name
	default:
		panic("can't get name for type of kind " + typ.kind.String())
	}
}

func (typ _type) Compare(other Type) bool {
	if other == nil {
		panic("can't compare: other should be non nil")
	}
	if typ.kind == Any {
		return true
	}
	if typ.kind != other.Kind() {
		return false
	}
	switch typ.kind {
	case Int, Float, String, Bool:
	case List, Null:
		return typ.subType.Compare(other.Subtype())
	case Function:
		if typ.NumIn() != other.NumIn() {
			return false
		}
		if !typ.ret.Compare(other.Out()) {
			return false
		}
		for i, tp := range typ.args {
			if !tp.Compare(other.In(i)) {
				return false
			}
		}
	case Record:
		if typ.name != other.Name() {
			return false
		}
		if typ.NumField() != other.NumField() {
			return false
		}
		for i, fld := range typ.fields {
			othFld := other.Field(i)
			if fld.Name() != othFld.Name() {
				return false
			}
			if !fld.Type().Compare(othFld.Type()) {
				return false
			}
		}
	default:
		panic("can't compare type of kind " + typ.kind.String())
	}
	return true
}

func (typ _type) Subtype() Type {
	switch typ.kind {
	case List, Null:
		return typ.subType
	default:
		panic("can't get subtype for " + typ.kind.String())
	}
}

func (typ _type) NumIn() int {
	if typ.kind != Function {
		panic("can't get number of arguments for " + typ.kind.String())
	}
	return len(typ.args)
}

func (typ _type) In(index int) Type {
	if typ.kind != Function {
		panic("can't get argument for " + typ.kind.String())
	}
	if index >= typ.NumIn() || index < 0 {
		panic("index out of bounds")
	}
	return typ.args[index]
}

func (typ _type) Out() Type {
	if typ.kind != Function {
		panic("can't get return type for " + typ.kind.String())
	}
	return typ.ret
}

func (typ _type) NumField() int {
	if typ.kind != Record {
		panic("can't get number of field for " + typ.kind.String())
	}
	return len(typ.fields)
}

func (typ _type) Field(index int) Field {
	if typ.kind != Record {
		panic("can't get field for " + typ.kind.String())
	}
	if index < 0 || index >= typ.NumField() {
		panic("index out of bounds")
	}
	return typ.fields[index]
}

func (typ _type) FieldIndex(name string) int {
	if typ.kind != Record {
		panic("can't get index of field for " + typ.kind.String())
	}
	for i, fld := range typ.fields {
		if fld.Name() == name {
			return i
		}
	}
	return -1
}

func (typ _type) FieldByName(name string) Field {
	if typ.kind != Record {
		panic("can't get index of field for " + typ.kind.String())
	}
	index := typ.FieldIndex(name)
	if index == -1 {
		return nil
	}
	return typ.fields[index]
}

func (typ _type) String() string {
	switch typ.kind {
	case Int, Float, String, Bool, Void, Any:
		return typ.kind.String()
	case List, Null:
		return fmt.Sprintf("%s<%s>", typ.kind.String(), typ.subType.Name())
	case Function:
		args := []string{}
		for _, arg := range typ.args {
			args = append(args, arg.Name())
		}
		return fmt.Sprintf("fn[%s]<%s>", strings.Join(args, " "), typ.ret.Name())
	case Record:
		fields := []string{}
		for _, field := range typ.fields {
			fields = append(fields, fmt.Sprintf("%s::%s", field.Name(), field.Type().Name()))
		}
		return fmt.Sprintf("%s<%s>", typ.name, strings.Join(fields, " "))
	default:
		panic("can't get string for type of kind " + typ.kind.String())
	}
}

type _field struct {
	name string
	tp   Type
}

func (fld _field) Name() string {
	return fld.name
}

func (fld _field) Type() Type {
	return fld.tp
}

func NewField(name string, typ Type) _field {
	return _field{
		name: name,
		tp:   typ,
	}
}

func NewIntType() _type {
	return _type{
		kind: Int,
	}
}

func NewFloatType() _type {
	return _type{
		kind: Float,
	}
}

func NewStringType() _type {
	return _type{
		kind: String,
	}
}

func NewBoolType() _type {
	return _type{
		kind: Bool,
	}
}

func NewAnyType() _type {
	return _type{
		kind: Any,
	}
}

func NewVoidType() Type {
	return _type{
		kind: Void,
	}
}

func NewNullType(sub Type) _type {
	return _type{
		kind:    Null,
		subType: sub,
	}
}

func NewListType(item Type) _type {
	return _type{
		kind:    List,
		subType: item,
	}
}

func NewFunctionType(ret Type, args ...Type) _type {
	if ret == nil {
		ret = NewVoidType()
	}
	return _type{
		kind: Function,
		args: args,
		ret:  ret,
	}
}

func NewRecordType(name string, fields ...Field) _type {
	if name == "" {
		panic("record name can't be empty")
	}
	return _type{
		kind:   Record,
		name:   name,
		fields: fields,
	}
}
