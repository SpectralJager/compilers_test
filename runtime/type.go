package runtime

import (
	"fmt"
	"strings"
)

type Type interface {
	Kind() Kind
	Name() string
	String() string
	Compare(Type) bool
	Item() Type
	NumIns() int
	In(int) Type
	Out() Type
	NumFields() int
	Field(string) (FieldType, int)
	FieldByIndex(int) FieldType
}

type typ struct {
	name   string
	kind   Kind
	item   Type
	args   []Type
	ret    Type
	fields []FieldType
}

func (typ *typ) Kind() Kind {
	return typ.kind
}

func (typ *typ) Name() string {
	switch typ.kind {
	case TY_Void:
		return "void"
	case TY_Int:
		return "int"
	case TY_Float:
		return "float"
	case TY_Bool:
		return "bool"
	case TY_String:
		return "string"
	case TY_List:
		return "list"
	case TY_Variatic:
		return "variatic"
	case TY_Function:
		return "fn"
	case TY_Record:
		return typ.name
	default:
		panic("can't get type name: unexpected kind of type")
	}
}

func (typ *typ) String() string {
	switch typ.kind {
	case TY_Int, TY_Float, TY_Bool, TY_String, TY_Void:
		return typ.Name()
	case TY_List:
		return fmt.Sprintf("%s<%s>", typ.Name(), typ.item.String())
	case TY_Variatic:
		return fmt.Sprintf("...%s", typ.item)
	case TY_Function:
		args := []string{}
		for _, arg := range typ.args {
			args = append(args, arg.String())
		}
		return fmt.Sprintf("%s[%s]<%s>", typ.Name(), strings.Join(args, " "), typ.ret.String())
	case TY_Record:
		fields := []string{}
		for _, fld := range typ.fields {
			fields = append(fields, fld.String())
		}
		return fmt.Sprintf("%s<%s>", typ.name, strings.Join(fields, " "))
	default:
		panic("can't get string of type: unexpected kind of type")
	}
}

func (typ *typ) Compare(other Type) bool {
	if typ.kind != other.Kind() {
		return false
	}
	switch typ.kind {
	case TY_Int, TY_Float, TY_Bool, TY_String, TY_Void:
	case TY_List, TY_Variatic:
		return other.Item().Compare(typ.item)
	case TY_Function:
		if typ.NumIns() != other.NumIns() || !other.Out().Compare(typ.ret) {
			return false
		}
		for i, in := range typ.args {
			if !in.Compare(other.In(i)) {
				return false
			}
		}
	case TY_Record:
		if typ.NumFields() != other.NumFields() || typ.Name() != other.Name() {
			return false
		}
		for i, fld := range typ.fields {
			oFld, j := other.Field(fld.Name())
			if !oFld.Type().Compare(fld.Type()) || i != j {
				return false
			}
		}
	}
	return true
}

func (typ *typ) Item() Type {
	if typ.kind != TY_List && typ.kind != TY_Variatic {
		panic("can't get item type: type is not list")
	}
	if typ.item == nil {
		panic("can't get item type: item type is empty")
	}
	return typ.item
}

func (typ *typ) NumIns() int {
	if typ.kind != TY_Function {
		panic("can't get number of input arguments: type is not function")
	}
	return len(typ.args)
}

func (typ *typ) In(index int) Type {
	if typ.kind != TY_Function {
		panic("can't get type of input argument: type is not function")
	}
	if index >= typ.NumIns() && index < 0 {
		panic("can't get type of input argument: index out of bounds")
	}
	return typ.args[index]
}

func (typ *typ) Out() Type {
	if typ.kind != TY_Function {
		panic("can't get type of output: type is not function")
	}
	if typ.ret == nil {
		return NewVoidType()
	}
	return typ.ret
}

func (typ *typ) NumFields() int {
	if typ.kind != TY_Record {
		panic("can't get number of fields: type is not record")
	}
	return len(typ.fields)
}

func (typ *typ) Field(name string) (FieldType, int) {
	if typ.kind != TY_Record {
		panic("can't get field: type is not record")
	}
	for i, fld := range typ.fields {
		if fld.Name() == name {
			return fld, i
		}
	}
	return nil, -1
}

func (typ *typ) FieldByIndex(index int) FieldType {
	if typ.kind != TY_Record {
		panic("can't get field: type is not record")
	}
	if index >= typ.NumFields() || index < 0 {
		panic("can't get field: index out of bounds")
	}
	return typ.fields[index]
}

func NewVoidType() *typ {
	return &typ{
		kind: TY_Void,
	}
}

func NewIntType() *typ {
	return &typ{
		kind: TY_Int,
	}
}

func NewFloatType() *typ {
	return &typ{
		kind: TY_Float,
	}
}
func NewStringType() *typ {
	return &typ{
		kind: TY_String,
	}
}
func NewBoolType() *typ {
	return &typ{
		kind: TY_Bool,
	}
}

func NewListType(item Type) *typ {
	return &typ{
		kind: TY_List,
		item: item,
	}
}

func NewVariaticType(item Type) *typ {
	return &typ{
		kind: TY_Variatic,
		item: item,
	}
}

func NewFunctionType(out Type, ins ...Type) *typ {
	return &typ{
		kind: TY_Function,
		args: ins,
		ret:  out,
	}
}

func NewRecordType(name string, fields ...FieldType) *typ {
	return &typ{
		kind:   TY_Record,
		name:   name,
		fields: fields,
	}
}

type FieldType interface {
	String() string
	Name() string
	Type() Type
}

type fieldType struct {
	name string
	typ  Type
}

func (f *fieldType) String() string {
	return fmt.Sprintf("%s::%s", f.name, f.typ.String())
}

func (f *fieldType) Name() string {
	return f.name
}

func (f *fieldType) Type() Type {
	return f.typ
}

func NewFieldType(name string, typ Type) *fieldType {
	return &fieldType{
		name: name,
		typ:  typ,
	}
}
